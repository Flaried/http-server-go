package server

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"

	"github.com/codecrafters-io/http-server-starter-go/app/internal/models"
	"github.com/codecrafters-io/http-server-starter-go/app/internal/router"
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

func (s *Server) Start() error {
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

func (s *Server) handleConnection(conn net.Conn) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	defer conn.Close()

	req, err := s.parseRequest(conn)
	if err != nil {
		return
	}

	s.router.Serve(conn, req)
}

func (s *Server) parseRequest(conn net.Conn) (models.Request, error) {
	var req models.Request
	reader := bufio.NewReader(conn)

	// Parse request line
	requestLine, err := reader.ReadString('\n')
	if err != nil {
		return req, err
	}

	parts := strings.Fields(requestLine)
	if len(parts) < 3 {
		return req, errors.New("invalid request line")
	}

	req.Method = parts[0]
	req.URL = parts[1]
	req.Path = strings.Split(parts[1], "/")

	// Parse headers
	headers := make(map[string]string)
	for {
		line, err := reader.ReadString('\n')
		if err != nil || line == "\r\n" {
			break
		}

		headerParts := strings.SplitN(line, ":", 2)
		if len(headerParts) == 2 {
			key := strings.TrimSpace(strings.ToLower(headerParts[0]))
			value := strings.TrimSpace(headerParts[1])
			headers[key] = value
		}
	}

	req.Headers = headers

	// Parse body if present
	if lengthStr := headers["content-length"]; lengthStr != "" {
		contentLength, err := strconv.Atoi(lengthStr)
		if err != nil {
			return req, err
		}

		body := make([]byte, contentLength)
		_, err = io.ReadFull(reader, body)
		if err != nil {
			return req, err
		}

		req.Body = body
	}

	return req, nil
}
