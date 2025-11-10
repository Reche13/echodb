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

func TestRESPValueGetMethods(t *testing.T) {
	t.Run("GetString", func(t *testing.T) {
		tests := []struct {
			name     string
			value    *RESPValue
			wantStr  string
			wantOk   bool
		}{
			{
				name:    "SimpleString type",
				value:   &RESPValue{Type: SimpleString, Str: "hello"},
				wantStr: "hello",
				wantOk:  true,
			},
			{
				name:    "BulkString type",
				value:   &RESPValue{Type: BulkString, Str: "world"},
				wantStr: "world",
				wantOk:  true,
			},
			{
				name:    "Error type",
				value:   &RESPValue{Type: Error, Str: "ERR message"},
				wantStr: "ERR message",
				wantOk:  true,
			},
			{
				name:    "Integer type",
				value:   &RESPValue{Type: Integer, Int: 123},
				wantStr: "",
				wantOk:  false,
			},
			{
				name:    "Array type",
				value:   &RESPValue{Type: Array, Array: []*RESPValue{}},
				wantStr: "",
				wantOk:  false,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, ok := tt.value.GetString()
				if got != tt.wantStr || ok != tt.wantOk {
					t.Errorf("GetString() = (%q, %v), want (%q, %v)", got, ok, tt.wantStr, tt.wantOk)
				}
			})
		}
	})

	t.Run("GetInteger", func(t *testing.T) {
		tests := []struct {
			name    string
			value   *RESPValue
			wantInt int64
			wantOk  bool
		}{
			{
				name:    "Integer type",
				value:   &RESPValue{Type: Integer, Int: 12345},
				wantInt: 12345,
				wantOk:  true,
			},
			{
				name:    "SimpleString type",
				value:   &RESPValue{Type: SimpleString, Str: "hello"},
				wantInt: 0,
				wantOk:  false,
			},
			{
				name:    "BulkString type",
				value:   &RESPValue{Type: BulkString, Str: "world"},
				wantInt: 0,
				wantOk:  false,
			},
			{
				name:    "Array type",
				value:   &RESPValue{Type: Array, Array: []*RESPValue{}},
				wantInt: 0,
				wantOk:  false,
			},
			{
				name:    "Error type",
				value:   &RESPValue{Type: Error, Str: "ERR"},
				wantInt: 0,
				wantOk:  false,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, ok := tt.value.GetInteger()
				if got != tt.wantInt || ok != tt.wantOk {
					t.Errorf("GetInteger() = (%d, %v), want (%d, %v)", got, ok, tt.wantInt, tt.wantOk)
				}
			})
		}
	})

	t.Run("GetArray", func(t *testing.T) {
		arr := []*RESPValue{{Type: SimpleString, Str: "test"}}
		tests := []struct {
			name    string
			value   *RESPValue
			wantArr []*RESPValue
			wantOk  bool
		}{
			{
				name:    "Array type",
				value:   &RESPValue{Type: Array, Array: arr},
				wantArr: arr,
				wantOk:  true,
			},
			{
				name:    "empty Array type",
				value:   &RESPValue{Type: Array, Array: []*RESPValue{}},
				wantArr: []*RESPValue{},
				wantOk:  true,
			},
			{
				name:    "SimpleString type",
				value:   &RESPValue{Type: SimpleString, Str: "hello"},
				wantArr: nil,
				wantOk:  false,
			},
			{
				name:    "Integer type",
				value:   &RESPValue{Type: Integer, Int: 123},
				wantArr: nil,
				wantOk:  false,
			},
			{
				name:    "BulkString type",
				value:   &RESPValue{Type: BulkString, Str: "world"},
				wantArr: nil,
				wantOk:  false,
			},
			{
				name:    "Error type",
				value:   &RESPValue{Type: Error, Str: "ERR"},
				wantArr: nil,
				wantOk:  false,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, ok := tt.value.GetArray()
				if !reflect.DeepEqual(got, tt.wantArr) || ok != tt.wantOk {
					t.Errorf("GetArray() = (%v, %v), want (%v, %v)", got, ok, tt.wantArr, tt.wantOk)
				}
			})
		}
	})
}

func TestConstructors(t *testing.T) {
	t.Run("NewSimpleString", func(t *testing.T) {
		val := NewSimpleString("test")
		if val.Type != SimpleString || val.Str != "test" {
			t.Errorf("NewSimpleString() = %+v, want Type=SimpleString, Str=test", val)
		}
	})

	t.Run("NewError", func(t *testing.T) {
		val := NewError("ERR message")
		if val.Type != Error || val.Str != "ERR message" {
			t.Errorf("NewError() = %+v, want Type=Error, Str=ERR message", val)
		}
	})

	t.Run("NewBulkString", func(t *testing.T) {
		val := NewBulkString("bulk")
		if val.Type != BulkString || val.Str != "bulk" {
			t.Errorf("NewBulkString() = %+v, want Type=BulkString, Str=bulk", val)
		}
	})

	t.Run("NewNullBulkString", func(t *testing.T) {
		val := NewNullBulkString()
		if val.Type != BulkString || val.Str != "" {
			t.Errorf("NewNullBulkString() = %+v, want Type=BulkString, Str=", val)
		}
	})

	t.Run("NewInteger", func(t *testing.T) {
		tests := []struct {
			name string
			n    int64
		}{
			{"positive", 123},
			{"negative", -456},
			{"zero", 0},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				val := NewInteger(tt.n)
				if val.Type != Integer || val.Int != tt.n {
					t.Errorf("NewInteger(%d) = %+v, want Type=Integer, Int=%d", tt.n, val, tt.n)
				}
			})
		}
	})

	t.Run("NewArray", func(t *testing.T) {
		t.Run("with elements", func(t *testing.T) {
			arr := []*RESPValue{NewSimpleString("a"), NewInteger(1)}
			val := NewArray(arr)
			if val.Type != Array || !reflect.DeepEqual(val.Array, arr) {
				t.Errorf("NewArray() = %+v, want Type=Array, Array=%v", val, arr)
			}
		})

		t.Run("empty array", func(t *testing.T) {
			val := NewArray([]*RESPValue{})
			if val.Type != Array || len(val.Array) != 0 {
				t.Errorf("NewArray([]) = %+v, want Type=Array, empty Array", val)
			}
		})

		t.Run("nil array", func(t *testing.T) {
			val := NewArray(nil)
			if val.Type != Array || val.Array != nil {
				t.Errorf("NewArray(nil) = %+v, want Type=Array, Array=nil", val)
			}
		})
	})
}