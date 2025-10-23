package store

import "time"

func (s *Store) IsExpired(val Value) bool {
	if val.ExpiresAt == 0 {
		return false
	}
	return time.Now().Unix() > val.ExpiresAt
}

func (s *Store) Expire(key string, expiresAt int64) bool {
	val, ok := s.GetValueOrExpire(key)
	if !ok {
		return false
	}

	s.mu.Lock()
	defer s.mu.Unlock()
	val.ExpiresAt = expiresAt
	s.data[key] = val

	return true
}

func (s *Store) Persist(key string) bool {
	val, ok := s.GetValueOrExpire(key)
	if !ok {
		return false
	}

	val.ExpiresAt = 0
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = val
	return true
}


func (s *Store) TTL(key string) int64 {
	val, ok := s.GetValueOrExpire(key)
	if !ok {
		return -2
	}

	if val.ExpiresAt == 0 {
		return -1
	}

	return val.ExpiresAt - time.Now().Unix()
}