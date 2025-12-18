package server

import (
	"fmt"
	"io"
	"log"
	"net"
	"sync"

	"github.com/reche13/echodb/internal/commands"
	"github.com/reche13/echodb/internal/config"
	"github.com/reche13/echodb/internal/protocol"
)

type Server struct {
	config *config.ServerConfig
	executor *commands.Executor
	clients map[net.Conn]*Client
	mu sync.RWMutex
	listener net.Listener
	wg sync.WaitGroup
}

type Client struct {
	conn net.Conn
	parser *protocol.Parser
	serializer *protocol.Serializer
}

func New(cfg *config.ServerConfig, executor *commands.Executor) *Server {
	return &Server{
		config: cfg,
		executor: executor,
		clients: make(map[net.Conn]*Client),
	}
}

func NewClient(conn net.Conn) *Client {
	return &Client{
		conn: conn,
		parser: protocol.NewParser(conn),
		serializer: protocol.NewSerializer(),
	}
}

func (s *Server) Start() error {
	PrintBanner()
	addr := fmt.Sprintf("%s:%d", s.config.Host, s.config.Port)

	l, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("cannot listen on %s: %w", addr, err)
	}
	defer l.Close()
	log.Println("Listening on ", addr)

	return s.AcceptConnections(l)
}

func (s *Server) Stop() {
	log.Println("shutting down server...")

	if s.listener != nil {
		s.listener.Close()
		s.listener = nil
	}

	s.mu.Lock()
	for conn := range s.clients {
		conn.Close()
		delete(s.clients, conn)
	}
	s.mu.Unlock()


	s.wg.Wait()
	log.Println("server stopped")
}

func (s *Server) AcceptConnections(l net.Listener) error {
	s.listener = l

	for {
		conn, err := l.Accept()
		if err != nil {
			if s.listener == nil {
				return nil
			}
			log.Println("failed to accept connection: ", err)
			continue
		}
		log.Printf("connected to client: %s", conn.RemoteAddr())

		s.wg.Add(1)
		go func(){
			defer s.wg.Done()
			s.handleConnection(conn)
		}()

	}
}

func (s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()

	client := NewClient(conn)

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
			if err == io.EOF {
        		log.Println("client disconnected:", client.conn.RemoteAddr())
        		return
    		}
    		if netErr, ok := err.(net.Error); ok && !netErr.Timeout() {
        		log.Println("client disconnected (network error):", client.conn.RemoteAddr(), err)
        		return
    		}
			log.Println("failed to parse:", err.Error())
			return
		}

		val := s.executor.Execute(args)

		out, err := client.serializer.Serialize(val)
		if err != nil {
			log.Println("failed to serialize:", err.Error())
			return
		}
		client.conn.Write(out)
	}
}