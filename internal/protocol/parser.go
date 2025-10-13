package protocol

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Parser struct {
	reader *bufio.Reader
}

func NewParser(r io.Reader) *Parser {
	return &Parser{reader: bufio.NewReader(r)}
}

func (p *Parser) Parse() (*RESPValue, error) {
	prefix, err := p.reader.ReadByte()
	if err != nil {
		return nil, err
	}

	switch prefix {
		case '+':
			return p.readSimpleString()
		case '-':
			return p.readError()
		case '*':
			return p.readArray()
		case '$':
			return p.readBulkString()
		case ':':
			return p.readInteger()
		default:
			return nil, fmt.Errorf("unknown prefix: %c", prefix)
	}
}

func (p *Parser) readSimpleString() (*RESPValue, error){
	str, err :=p.reader.ReadString('\n')
	if err != nil {
		return nil, err
	}
	str = strings.TrimSuffix(str, "\r\n")
	return &RESPValue{
		Type: SimpleString,
		Str: str,
	}, nil
}

func (p *Parser) readError() (*RESPValue, error){
	str, err :=p.reader.ReadString('\n')
	if err != nil {
		return nil, err
	}
	str = strings.TrimSuffix(str, "\r\n")
	return &RESPValue{
		Type: Error,
		Str: str,
	}, nil	
}

func (p *Parser) readArray()(*RESPValue, error){
	line, err := p.reader.ReadString('\n')
	if err != nil {
		return nil, err
	}
	line = strings.TrimSuffix(line, "\r\n")
	arrLen, err := strconv.Atoi(line)
	if err != nil {
		return nil, err
	}

	elements := make([]*RESPValue, 0, arrLen)
	for i:= 0; i < arrLen; i++ {
		elem, err := p.Parse()
		if err != nil {
			return nil, err
		}
		elements = append(elements, elem)
	}
	return &RESPValue{
		Type: Array,
		Array: elements,
	}, nil
}

func (p *Parser) readBulkString() (*RESPValue, error){
	line, err := p.reader.ReadString('\n')
	if err != nil {
    	return nil, err
	}
	line = strings.TrimSuffix(line, "\r\n")
	readLen, err := strconv.Atoi(line)
	if err != nil {
		return nil, err
	}

	if readLen == -1 {
    	return nil, nil
	}

	buf := make([]byte, readLen+2)
	_, err = p.reader.Read(buf)
	if err != nil {
		return nil, err
	}

	str := string(buf[:readLen])
	return &RESPValue{
		Type: BulkString,
		Str: str,
	}, nil
}

func (p *Parser) readInteger() (*RESPValue, error){
	line, err := p.reader.ReadString('\n')
	if err != nil {
		return nil, err
	}

	line = strings.TrimSuffix(line, "\r\n")
	n, err := strconv.ParseInt(line, 10, 64)
	if err != nil {
		return nil, err
	}
	return &RESPValue{
		Type: Integer,
		Int: n,
	}, nil
}