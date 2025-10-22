package store

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/reche13/echodb/internal/protocol"
)

type AOFManager struct {
	file *os.File
	writer *bufio.Writer
	mu sync.Mutex
}

func NewAOFManager(path string) (*AOFManager, error) {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}

	return &AOFManager{
		file: file,
		writer: bufio.NewWriter(file),
	}, nil
}

func (a *AOFManager) AppendRESP(cmd string, args ...string) error {
	a.mu.Lock()
	defer a.mu.Unlock()

	n := 1 + len(args)
	data := fmt.Sprintf("*%d\r\n", n)

	data += fmt.Sprintf("$%d\r\n%s\r\n", len(cmd), cmd)

	for _, arg := range args {
		data += fmt.Sprintf("$%d\r\n%s\r\n", len(arg), arg)
	}

	if _, err := a.writer.WriteString(data); err != nil {
		return err
	}

	return a.writer.Flush()
}

func (a *AOFManager) Close() error {
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.writer != nil {
		if err := a.writer.Flush(); err != nil {
			return err
		}
	}

	if a.file != nil {
		return a.file.Close()
	}

	return nil
}


func (a *AOFManager) LoadFromAOF(s *Store) error {
	file, err := os.Open(a.file.Name())
	if err != nil {
		return err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	parser := protocol.NewParser(reader)

	for {
		respVal, err := parser.Parse()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return fmt.Errorf("failed to parse AOF: %w", err)
		}

		if respVal.Type != protocol.Array || len(respVal.Array) == 0 {
			continue
		}

		cmdStr, ok := respVal.Array[0].GetString()
		if !ok {
			continue
		}

		cmd := strings.ToUpper(cmdStr)
		args := []string{}
		for _, v := range respVal.Array[1:] {
			arg, ok := v.GetString()
			if !ok {
				continue
			}
			args = append(args,arg)
		}

		switch cmd {
		case "SET":
			if len(args) != 2 {
				continue
			}
			s.Set(args[0], args[1])
		case "DEL":
			if len(args) < 1 {
				continue
			}
			s.Del(args...)
		case "LPUSH":
			if len(args) < 2 {
				continue
			}
			for _, val := range args[1:] {
				s.LPush(args[0], val)
			}
		case "RPUSH":
			if len(args) < 2 {
				continue
			}
			for _, val := range args[1:] {
				s.RPush(args[0], val)
			}
		case "EXPIRE":
			if len(args) != 2 {
				continue
			}
			ttl, err := strconv.Atoi(args[1])
			if err != nil {
				continue
			}
			s.Expire(args[0], int64(ttl))
		case "PERSIST":
			if len(args) != 1 {
				continue
			}
			s.Persist(args[0])
		default:
			
		}
	}

	return nil
}
