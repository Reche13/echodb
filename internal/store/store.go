package store

import "sync"

type Store struct {
	data map[string]string
	mu sync.RWMutex
}

func New() *Store {
	return &Store{
		data: make(map[string]string),
	}
}

func (s *Store) Set(key, value string) {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.data[key] = value
}

func (s *Store) Get(key string) (string, bool) {
    s.mu.RLock()
    defer s.mu.RUnlock()
    val, ok := s.data[key]
    return val, ok
}

func (s *Store) Del(keys ...string) int {
    s.mu.Lock()
    defer s.mu.Unlock()

    deleted := 0
    for _, key := range keys {
        if _, ok := s.data[key]; ok {
            delete(s.data, key)
            deleted++
        }
    }
    return deleted
}

func (s *Store) Exists(keys ...string) int {
    s.mu.RLock()
    defer s.mu.RUnlock()

    count := 0
    for _, key := range keys {
        if _, ok := s.data[key]; ok {
            count++
        }
    }
    return count
}