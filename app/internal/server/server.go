package server

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"

	"github.com/codecrafters-io/http-server-starter-go/app/internal/config"
)

type Server struct {
	Listener net.Listener
	Router   *Router
	config   config.ServerConfig
}

func (s *Server) Start(config config.ServerConfig) {
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

		if strings.Contains(builder.String(), config.MARKER) {
			break
		}
	}

	requestString := builder.String()
	return requestString
}

// and then make all the request funcs use it
func (s *Server) Read(conn net.Conn) config.Request {
	connString := connectionToString(conn)

	var request config.Request
	var headers = make(map[string]string)
	iter := strings.SplitSeq(connString, config.CRLF)
	for partString := range iter {
		parts := strings.Fields(partString)

		if len(parts) == 0 {
			continue
		}

		switch strings.ToLower(parts[0]) {
		case "get":
			if len(parts) >= 3 {
				request.Method = parts[0]
				request.URL = parts[1]
				request.Body = parts[2]
			} else if len(parts) == 2 {
				request.Method = parts[0]
				request.URL = parts[1]
			}
		default:
			headerParts := strings.SplitN(partString, ":", 2)
			if len(headerParts) == 2 {
				key := strings.TrimSpace(headerParts[0])
				value := strings.TrimSpace(headerParts[1])
				headers[key] = value
			}
		}
	}

	request.Headers = headers
	request.UrlParts = strings.Split(request.URL, "/")
	return request

}

func (s *Server) handlerConnection(conn net.Conn) {
	defer conn.Close()
	read := s.Read(conn)
	s.Router.Serve(conn, &read)
}
