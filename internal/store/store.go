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

func (s *Store) TypeOf(key string) ValueType {
    s.mu.RLock()
    defer s.mu.RUnlock()

    val, ok := s.data[key]
    if !ok {
        return ""
    }

    return val.Type
}

func (s *Store) GetValue(key string) (Value, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	val, ok := s.data[key]
	if !ok {
		return Value{}, false
	}
	return val, true
}

func (s *Store) GetValueOrExpire(key string) (Value, bool) {
	val, ok := s.GetValue(key)
	if !ok {
		return Value{}, false
	}

	if s.IsExpired(val) {
        s.DeleteKey(key)
        return Value{}, false
    }

    return val, true
}

func (s *Store) DeleteKey(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.data, key)
}