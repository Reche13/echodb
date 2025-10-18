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

func NewBulkString(s string) *RESPValue {
	return &RESPValue{
		Type: BulkString,
		Str: s,
	}
}

func NewNullBulkString() *RESPValue {
	return &RESPValue{
		Type: BulkString,
		Str: "",
	}
}

func NewInteger(n int64) *RESPValue {
	return &RESPValue{
		Type: Integer,
		Int: n,
	}
}

func NewArray(arr []*RESPValue) *RESPValue {
	return &RESPValue{
		Type: Array,
		Array: arr,
	}
}