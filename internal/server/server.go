package server

import (
	"fmt"
	"log"
	"net"

	"github.com/reche13/echodb/internal/protocol"
)

type Server struct {
	Addr string
}

func New(addr string) *Server {
	return &Server{Addr: addr}
}

func (s *Server) Start() error {
	l, err := net.Listen("tcp", s.Addr)
	if err != nil {
		return fmt.Errorf("cannot listen on %s: %w", s.Addr, err)
	}
	defer l.Close()
	log.Println("Listening on ", s.Addr)

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println("failed to accept connection: ", err)
			continue
		}
		log.Printf("connected to client: %s", conn.RemoteAddr())

		go s.handleConnection(conn)
	}
}


func (s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()
	for {
		p := protocol.NewParser(conn)
		val, err := p.Parse()
		if err != nil {
			log.Println("failed to parse:", err)
			return
		}

		sr := protocol.NewSerializer()
		out, err := sr.Serialize(val)
		if err != nil {
			log.Println("failed to serialize:", err)
			return
		}

		fmt.Printf("%+v\n", val)
		conn.Write(out)
	}
}