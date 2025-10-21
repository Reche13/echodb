package store

import (
	"bufio"
	"fmt"
	"os"
	"sync"
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

