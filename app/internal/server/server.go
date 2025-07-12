package server

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"strings"

	"github.com/codecrafters-io/http-server-starter-go/app/internal/constants"
)

type Server struct {
	Listener         net.Listener
	Router           *Router
	config           constants.ServerConfig
	ServingDirectory *string
}

func (s *Server) Start(config constants.ServerConfig) {
	s.config = config
	s.Listen()

	defer s.Close()
	for {
		conn := s.Accept()
		go s.handlerConnection(conn)
	}

}

func (s *Server) Listen() {
	l, err := net.Listen(s.config.Protocol, s.config.Address)
	if err != nil {
		fmt.Printf("Failed to bind to %s err: %v", s.config, err)
		os.Exit(1)
	}

	s.Listener = l
}

func (s *Server) Accept() net.Conn {
	conn, err := s.Listener.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}
	return conn
}

func (s *Server) Close() {
	err := s.Listener.Close()
	if err != nil {
		fmt.Println("Failed to close listener:", err.Error())
	}

}

func connectionToString(conn net.Conn) string {
	reader := bufio.NewReader(conn)
	var builder strings.Builder

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				break
			}
			fmt.Printf("Error reading %v\n", err)
			break
		}

		builder.WriteString(line)

		if strings.Contains(builder.String(), constants.MARKER) {
			break
		}
	}

	requestString := builder.String()
	return requestString
}

// and then make all the request funcs use it
func (s *Server) Read(conn net.Conn) (Request, error) {
	var request Request
	reader := bufio.NewReader(conn)

	requestLine, err := reader.ReadString('\n')
	if err != nil {
		return request, err
	}
	requestParts := strings.Fields(requestLine)
	if len(requestParts) < 3 {
		return request, errors.New("bad request")
	}

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

	request.Headers = headers
	request.Method = requestParts[0]
	request.Path = strings.Split(requestParts[1], "/")
	return request, nil

}

func (s *Server) handlerConnection(conn net.Conn) {
	// Recover from panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic in connection handler:", r)

		}
	}()

	defer conn.Close()

	request, err := s.Read(conn)
	if err == nil {
		s.Router.Serve(conn, &request)
	}
}
