package persistence

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/reche13/echodb/internal/config"
	"github.com/reche13/echodb/internal/protocol"
	"github.com/reche13/echodb/internal/store"
)

type AOFManager struct {
	file *os.File
	config *config.PersistenceConfig
	writer *bufio.Writer
	syncMode SyncMode
	mu sync.Mutex
}

type SyncMode string

const  (
	ALWAYS SyncMode = "always"
	EVERYSEC SyncMode = "everysec"
	NO SyncMode = "no" 
)

func NewAOFManager(cfg *config.PersistenceConfig) (*AOFManager, error) {
	path := filepath.Join(cfg.DataDir, cfg.Aof.File)
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}

	return &AOFManager{
		file:   file,
		writer: bufio.NewWriter(file),
		config: cfg,
		syncMode: SyncMode(cfg.Aof.SyncMode),
	}, nil
}


func (a *AOFManager) Log(cmd *protocol.RESPValue) error {
	a.mu.Lock()
	defer a.mu.Unlock()

	sr := protocol.NewSerializer()
	data, err := sr.Serialize(cmd)
	if err != nil {
		return err
	}

	if _, err := a.writer.Write(data); err != nil {
		return err
	}

	switch a.syncMode {
	case ALWAYS:
		return a.writer.Flush()
	case EVERYSEC:
		return nil
	case NO:
		return nil
	default:
		return a.writer.Flush()
	}
}

func (a *AOFManager) StartBackgroundFlush() {
	if a.syncMode != EVERYSEC {
		return
	}

	go func(){
		ticker := time.NewTicker(10 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				a.mu.Lock()
				a.writer.Flush()
				a.mu.Unlock()
			}
		}
	}()
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


func (a *AOFManager) Load(s *store.Store) error {
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
		case "LPOP":
			if len(args) < 1 {
				continue
			}
			if len(args) == 2 {
				count, err := strconv.Atoi(args[1])
				if err != nil {
					continue
				}
				s.LPop(args[0], count)
			} else {
				s.LPop(args[0], 1)
			}
		case "RPOP":
			if len(args) < 1 {
				continue
			}
			if len(args) == 2 {
				count, err := strconv.Atoi(args[1])
				if err != nil {
					continue
				}
				s.LPop(args[0], count)
			} else {
				s.LPop(args[0], 1)
			}
		case "EXPIRE":
			if len(args) != 2 {
				continue
			}
			ttl, err := strconv.Atoi(args[1])
			if err != nil {
				continue
			}
			s.Expire(args[0], time.Now().Unix() + int64(ttl))
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
