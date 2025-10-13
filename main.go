package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"

	"github.com/reche13/echodb/internal/protocol"
)

type ValueType string

const (
	ARRAY ValueType = "*"
	BULK ValueType = "$"
	STRING ValueType = "+"
)

type Value struct {
	typ ValueType
	bulk string
	str string
	array []Value
}

func (v *Value) readArray(reader io.Reader) {
	buf := make([]byte, 4)
	_, err := reader.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	arrLen, err := strconv.Atoi(string(buf[1]))
	if err != nil {
		fmt.Println(err)
		return
	}

	for range arrLen {
		bulk := v.readBulk(reader)
		v.array = append(v.array, bulk)
	}
}

func (v *Value) readBulk(reader io.Reader) Value {
	buf := make([]byte, 2)
	_, err := reader.Read(buf)
	if err != nil {
		fmt.Println(err)
		return Value{}
	}

	n, err := strconv.Atoi(string(buf[1]))
	if err != nil {
		fmt.Println(err)
		return Value{}
	}

	bulkBuff := make([]byte, n+2)
	_, err = reader.Read(bulkBuff)
	if err != nil {
		fmt.Println(err)
		return Value{}
	}

	bulk := string(bulkBuff[:n])
	return Value{typ: BULK, bulk: bulk}
}

func main() {
	l, err := net.Listen("tcp", ":6380")
	if err != nil {
		log.Fatal("cannot listen on port :6380")
	}
	log.Println("Listening on port :6380")

	defer l.Close()

	conn, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer conn.Close()
	log.Printf("connected to client: %s", conn.RemoteAddr())

	for {
		p := protocol.NewParser(conn)
		val, err := p.Parse()
		if err != nil {
			log.Fatal("failed to parse")
			break
		}
		
		fmt.Println(val)

		conn.Write([]byte("+OK\r\n"))
	}
}