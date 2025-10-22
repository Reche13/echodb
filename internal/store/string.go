package store

func (s *Store) Set(key, value string) {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.data[key] = Value{Type: StringType, Data: value}

    if s.Aof != nil {
        s.Aof.AppendRESP("SET", key, value)
    }
}

func (s *Store) Get(key string) (string, bool) {
    val, ok := s.GetValueOrExpire(key)
    if !ok || val.Type != StringType {
        return "", false
    }

    return val.Data.(string), ok
}

func (s *Store) Del(keys ...string) int {
    deleted := 0
    for _, key := range keys {
        if _, ok := s.GetValueOrExpire(key); ok {
            s.DeleteKey(key)
            if s.Aof != nil {
                s.Aof.AppendRESP("DEL", key)
            }
            deleted++
        }
    }
    return deleted
}

func (s *Store) Exists(keys ...string) int {
    count := 0
    for _, key := range keys {
        if _, ok := s.GetValueOrExpire(key); ok {
            count++
        }
    }
    return count
}