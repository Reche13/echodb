package store

func (s *Store) TypeOf(key string) ValueType {
    s.mu.RLock()
    defer s.mu.RUnlock()

    val, ok := s.data[key]
    if !ok {
        return ""
    }

    return val.Type
}

func (s *Store) Set(key, value string) {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.data[key] = Value{Type: StringType, Data: value}
}

func (s *Store) Get(key string) (string, bool) {
    s.mu.RLock()
    defer s.mu.RUnlock()
    val, ok := s.data[key]
    if !ok || val.Type != StringType {
        return "", false
    }
    return val.Data.(string), ok
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