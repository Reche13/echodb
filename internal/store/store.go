package store

import (
	"sync"
)

type ValueType string

const (
    StringType ValueType = "string"
    ListType ValueType = "list"
)

type Value struct {
    Type ValueType
    Data any
    ExpiresAt int64
}

type Store struct {
	data map[string]Value
	mu sync.RWMutex
}

func New() *Store {
	return &Store{
		data: make(map[string]Value),
	}
}