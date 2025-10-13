package protocol

import (
	"bytes"
	"fmt"
)

type Serializer struct {
	buffer *bytes.Buffer
}

func NewSerializer() *Serializer {
	return &Serializer{
		buffer: &bytes.Buffer{},
	}
}

func (s *Serializer) Serialize(rv *RESPValue) ([]byte, error) {
	s.buffer.Reset()	
	if err := s.serializeValue(rv); err != nil {
		return nil, err
	}
	return s.buffer.Bytes(), nil
}

func (s *Serializer) serializeValue(rv *RESPValue) error {
	switch rv.Type {
		case SimpleString:
			return s.writeSimpleString(rv.Str)
		case Error:
			return s.writeError(rv.Str)
		case Integer:
			return s.writeInteger(rv.Int)
		case BulkString:
			return s.writeBulkString(rv.Str)
		case Array:
			return s.writeArray(rv.Array)
		default:
			return fmt.Errorf("unknown RESP type: %v", rv.Type)
	}
}

func (s *Serializer) writeSimpleString(v string) error {
	if v == "" {
		_, err := s.buffer.WriteString("$-1\r\n")
		if err != nil {
			return err
		}
	}
	_, err := fmt.Fprintf(s.buffer, "+%s\r\n", v)
	return err
}

func (s *Serializer) writeError(v string) error {
	_, err := fmt.Fprintf(s.buffer, "-%s\r\n", v)
	return err
}

func (s *Serializer) writeInteger(v int64) error {
	_, err := fmt.Fprintf(s.buffer, ":%d\r\n", v)
	return err
}

func (s *Serializer) writeBulkString(v string) error {
	_, err := fmt.Fprintf(s.buffer, "$%d\r\n%s\r\n", len(v), v)
	return err
}

func (s *Serializer) writeArray(arr []*RESPValue) error {
	if arr == nil {
		_, err := s.buffer.WriteString("*-1\r\n")
		return err
	}

	_, err := fmt.Fprintf(s.buffer, "*%d\r\n", len(arr))
	if err != nil {
		return err
	}

	for _, val := range arr {
		if err := s.serializeValue(val); err != nil {
			return err
		}
	}

	return nil
}