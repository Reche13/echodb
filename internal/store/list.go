package store

func (s *Store) LPush(key string, values ...string) int {
	s.mu.Lock()
	defer s.mu.Unlock()

	val, ok := s.data[key]
	var list []string
	if ok {
		if val.Type != ListType {
			return -1
		}
		list = val.Data.([]string)
	}

	for i := len(values)/2 - 1; i >= 0; i-- {
		opp := len(values) -1 - i
		values[i], values[opp] = values[opp], values[i]
	}

	list = append(values, list...)
	s.data[key] = Value{Type: ListType, Data: list, ExpiresAt: val.ExpiresAt}

	return len(list)
}

func (s *Store) LPop(key string, count int) []string {
	s.mu.Lock()
	defer s.mu.Unlock()

	if count <= 0 {
		return nil
	}

	val, ok := s.data[key]
	if !ok || val.Type != ListType {
		return nil
	}

	list := val.Data.([]string)
	if len(list) == 0 {
		return nil
	}

	if count > len(list) {
		count = len(list)
	}

	popped := list[:count]
	list = list[count:]
	s.data[key] = Value{Type: ListType, Data: list,ExpiresAt: val.ExpiresAt}
	
	return popped
}

func (s *Store) LRange(key string, start int, stop int) []string {
	s.mu.RLock()
	defer s.mu.RUnlock()

	val, ok := s.data[key]
	if !ok || val.Type != ListType {
		return []string{}
	}

	list := val.Data.([]string)
	length:= len(list)

	if length == 0 {
		return []string{}
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
		return []string{}
	}

	sublist:= list[start : stop+1]

	return sublist
}


func (s *Store) RPush(key string, values ...string) int {
	s.mu.Lock()
	defer s.mu.Unlock()

	val, ok := s.data[key]
	var list []string
	if ok {
		if val.Type != ListType {
			return -1
		}
		list = val.Data.([]string)
	}

	list = append(list, values...)
	s.data[key] = Value{Type: ListType, Data: list, ExpiresAt: val.ExpiresAt}

	return len(list)
}

func (s *Store) RPop(key string, count int) []string {
	s.mu.Lock()
	defer s.mu.Unlock()

	if count <= 0 {
		return nil
	}

	val, ok := s.data[key]
	if !ok || val.Type != ListType {
		return nil
	}

	list := val.Data.([]string)
	if len(list) == 0 {
		return nil
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
	s.data[key] = Value{Type: ListType, Data: list,ExpiresAt: val.ExpiresAt}
	
	return popped
}