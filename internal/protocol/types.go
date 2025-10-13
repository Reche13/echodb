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