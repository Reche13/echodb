package store

import "fmt"

func (s *Store) LPush(key string, values ...string) (int, error) {
	val, ok := s.GetValueOrExpire(key)
	var list []string
	expiresAt := int64(0)
	if ok {
		if val.Type != ListType {
			return 0, fmt.Errorf("WRONGTYPE Operation against a key holding the wrong kind of value")
		}
		list = val.Data.([]string)
		expiresAt = val.ExpiresAt
	}

	for i := len(values)/2 - 1; i >= 0; i-- {
		opp := len(values) -1 - i
		values[i], values[opp] = values[opp], values[i]
	}

	list = append(values, list...)

	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = Value{Type: ListType, Data: list, ExpiresAt: expiresAt}

	return len(list), nil
}

func (s *Store) LPop(key string, count int) ([]string, error) {
	val, ok := s.GetValueOrExpire(key)
	if !ok {
		return []string{}, nil
	}

	if val.Type != ListType {
		return nil, fmt.Errorf("WRONGTYPE Operation against a key holding the wrong kind of value")
	}

	if count <= 0 {
		return []string{}, nil
	}

	list := val.Data.([]string)
	if len(list) == 0 {
		return []string{}, nil
	}

	if count > len(list) {
		count = len(list)
	}

	popped := list[:count]
	list = list[count:]

	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = Value{Type: ListType, Data: list,ExpiresAt: val.ExpiresAt}
	
	return popped, nil
}

func (s *Store) LRange(key string, start int, stop int) ([]string, error) {
	val, ok := s.GetValueOrExpire(key)
	if !ok {
		return []string{}, nil
	}

	if val.Type != ListType {
		return nil, fmt.Errorf("WRONGTYPE Operation against a key holding the wrong kind of value")
	}

	list := val.Data.([]string)
	length:= len(list)

	if length == 0 {
		return []string{}, nil
	}

	if start < 0 {
		start = length + start
	}

	if stop < 0 {
		stop = length + stop
	}

	if start < 0 {
		start = 0
	}

	if stop >= length {
		stop = length - 1
	}

	if start > stop || start >= length {
		return []string{}, nil
	}

	sublist:= list[start : stop+1]

	return sublist, nil
}

func (s *Store) RPush(key string, values ...string) (int, error) {
	val, ok := s.GetValueOrExpire(key)
	var list []string
	expiresAt := int64(0)
	if ok {
		if val.Type != ListType {
			return 0, fmt.Errorf("WRONGTYPE Operation against a key holding the wrong kind of value")
		}
		list = val.Data.([]string)
		expiresAt = val.ExpiresAt
	}

	list = append(list, values...)

	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = Value{Type: ListType, Data: list, ExpiresAt: expiresAt}

	return len(list), nil
}

func (s *Store) RPop(key string, count int) ([]string, error) {
	val, ok := s.GetValueOrExpire(key)
	if !ok {
		return []string{}, nil
	}

	if val.Type != ListType {
		return nil, fmt.Errorf("WRONGTYPE Operation against a key holding the wrong kind of value")
	}

	if count <= 0 {
		return []string{}, nil
	}

	list := val.Data.([]string)
	if len(list) == 0 {
		return []string{}, nil
	}

	if count > len(list) {
		count = len(list)
	}
	start := len(list) - count

	popped := make([]string, count)
	for i := 0; i < count; i++ {
		popped[i] = list[len(list)-1-i]
	}

	list = list[:start]

	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = Value{Type: ListType, Data: list,ExpiresAt: val.ExpiresAt}
	
	return popped, nil
}

func (s *Store) LLen(key string) (int, error) {
	val, ok := s.GetValueOrExpire(key)
	if !ok {
		return 0, nil
	}
	if val.Type != ListType {
		return 0, fmt.Errorf("WRONGTYPE Operation against a key holding the wrong kind of value")
	}
	list, _ := val.Data.([]string)
	return len(list), nil
}