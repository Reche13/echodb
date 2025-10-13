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

func (p *Parser) Parse() (any, error) {
	prefix, err := p.reader.ReadByte()
	if err != nil {
		return nil, err
	}

	switch prefix {
		case '+':
			return p.readSimpleString()
		case '-':
			return "ReadError", nil
		case '$':
			return "ReadBulkString", nil
		case ':':
			return "ReadInteger", nil
		case '*':
			return "ReadArray", nil
		default:
			return nil, fmt.Errorf("unknown prefix: %c", prefix)
	}
}

func (p *Parser) readSimpleString() (string, error){
	str, err :=p.reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	str = strings.TrimSuffix(str, "\r\n")
	return str, err
}

func (p *Parser) readError() (string, error){
	str, err :=p.reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	str = strings.TrimSuffix(str, "\r\n")
	return str, err
}

func (p *Parser) readBulkString() (string, error){
	line, err := p.reader.ReadString('\n')
	if err != nil {
    	return "", err
	}
	line = strings.TrimSuffix(line, "\r\n")
	readLen, err := strconv.Atoi(line)

	if readLen == -1 {
    	return "", nil
	}

	buf := make([]byte, readLen+2)
	_, err = p.reader.Read(buf)
	if err != nil {
		return "", err
	}

	str := string(buf[:readLen])
	return str, nil
}

func (p *Parser) readArray()(any, error){
	return nil, nil
}