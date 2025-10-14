package protocol

type RESPType int

const (
	SimpleString RESPType = iota
	Error
	Array
	BulkString
	Integer
)

type RESPValue struct {
	Type  RESPType
	Str   string
	Int   int64
	Array []*RESPValue
}

func (rv *RESPValue) GetString() (string, bool) {
	if rv.Type == SimpleString || rv.Type == BulkString || rv.Type == Error {
		return rv.Str, true
	}
	return "", false
}

func (rv *RESPValue) GetInteger() (int64, bool) {
	if rv.Type == Integer {
		return rv.Int, true
	}
	return 0, false
}

func (rv *RESPValue) GetArray() ([]*RESPValue, bool) {
	if rv.Type == Array {
		return rv.Array, true
	}
	return nil, false
}