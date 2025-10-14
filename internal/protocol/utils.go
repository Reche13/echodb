package protocol

func NewSimpleString(s string) *RESPValue {
	return &RESPValue{
		Type: SimpleString,
		Str: s,
	}
}

func NewError(s string) *RESPValue {
	return &RESPValue{
		Type: Error,
		Str: s,
	}
}

