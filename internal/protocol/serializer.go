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
		default:
			return nil
	}
}

func (s *Serializer) writeSimpleString(v string) error {
	_, err := fmt.Fprintf(s.buffer, "+%s\r\n", v)
	return err
}