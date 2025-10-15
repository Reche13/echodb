package server

import (
	"fmt"
	"log"
	"net"
	"sync"

	"github.com/reche13/echodb/internal/commands"
	"github.com/reche13/echodb/internal/protocol"
	"github.com/reche13/echodb/internal/store"
)

type Server struct {
	Addr string
	store *store.Store
	clients map[net.Conn]*Client
	mu sync.RWMutex
}

type Client struct {
	conn net.Conn
	parser *protocol.Parser
	serializer *protocol.Serializer
}

func New(addr string, store *store.Store) *Server {
	return &Server{
		Addr: addr,
		store: store,
	}
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

	client := &Client{
		conn: conn,
		parser: protocol.NewParser(conn),
		serializer: protocol.NewSerializer(),
	}

	s.mu.Lock()
	s.clients[conn] = client
	s.mu.Unlock()

	defer func(){
		s.mu.Lock()
		delete(s.clients, conn)
		s.mu.Unlock()
	}()

	for {
		args, err := client.parser.Parse()
		if err != nil {
			log.Println("failed to parse:", err)
			return
		}

		val := commands.Execute(s.store ,args)

		out, err := client.serializer.Serialize(val)
		if err != nil {
			log.Println("failed to serialize:", err)
			return
		}
		client.conn.Write(out)
	}
}