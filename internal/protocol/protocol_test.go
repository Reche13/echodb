package protocol

import (
	"bytes"
	"reflect"
	"testing"
)

func TestParserAndSerializer(t *testing.T) {
	tests := []struct {
		name string
		input []byte
		want *RESPValue
	}{
		{
			name:  "simple string",
			input: []byte("+hello\r\n"),
			want:  &RESPValue{Type: SimpleString, Str: "hello"},
		},
		{
			name:  "error",
			input: []byte("-ERR something went wrong\r\n"),
			want:  &RESPValue{Type: Error, Str: "ERR something went wrong"},
		},
		{
			name:  "integer",
			input: []byte(":12345\r\n"),
			want:  &RESPValue{Type: Integer, Int: 12345},
		},
		{
			name:  "bulk string",
			input: []byte("$5\r\nhello\r\n"),
			want:  &RESPValue{Type: BulkString, Str: "hello"},
		},
		{
			name: "array of bulk strings",
			input: []byte("*3\r\n$3\r\nset\r\n$1\r\na\r\n$5\r\nhello\r\n"),
			want: &RESPValue{
				Type: Array,
				Array: []*RESPValue{
					{Type: BulkString, Str: "set"},
					{Type: BulkString, Str: "a"},
					{Type: BulkString, Str: "hello"},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parser := NewParser(bytes.NewReader(tt.input))
			got, err := parser.Parse()
			if err != nil {
				t.Fatalf("Parse() error = %v", err)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() = %+v, want %+v", got, tt.want)
			}

			serializer := NewSerializer()
			out, err := serializer.Serialize(got)
			if err != nil {
				t.Fatalf("Serialize() error = %v", err)
			}

			if !bytes.Equal(out, tt.input) {
				t.Errorf("Serialize() = %q, want %q", out, tt.input)
			}
		})
	}
}