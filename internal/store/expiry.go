package store

import "time"

func (s *Store) IsExpired(val Value) bool {
	if val.ExpiresAt == 0 {
		return false
	}
	return time.Now().Unix() > val.ExpiresAt
}