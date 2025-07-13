package server

import (
	"fmt"
	"github.com/codecrafters-io/http-server-starter-go/app/internal/models"
	"github.com/codecrafters-io/http-server-starter-go/app/internal/router"
	"net"
	"strings"
)

type Server struct {
	listener net.Listener
	router   *router.Router
	config   models.ServerConfig
}

func NewServer(config models.ServerConfig) *Server {
	return &Server{
		config: config,
		router: router.NewRouter(),
	}
}

func (s *Server) SetRouter(r *router.Router) {
	s.router = r
}

func (s Server) Start() error {
	if err := s.listen(); err != nil {
		return err
	}

	defer s.Close()

	for {
		conn, err := s.listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go s.handleConnection(conn)
	}
}

func (s *Server) listen() error {
	l, err := net.Listen(s.config.Protocol, s.config.Address)
	if err != nil {
		return fmt.Errorf("failed to bind to %s: %w", s.config.Address, err)
	}

	s.listener = l
	return nil
}

func (s *Server) Close() {
	if s.listener != nil {
		s.listener.Close()
	}
}

// main logic for incoming requests on a single client connection.
func (s Server) handleConnection(conn net.Conn) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	defer conn.Close()

	// Loop to handle multiple requests on same connection
	for {
		req, err := models.ParseRequest(conn)
		if err != nil {
			return
		}

		// Check if client wants to close connection
		if strings.ToLower(req.Headers["connection"]) == "close" {
			fmt.Println("close.")
			s.router.Serve(conn, req)
			return
		}

		s.router.Serve(conn, req)
	}
}
