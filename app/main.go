package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"

	"github.com/codecrafters-io/http-server-starter-go/app/models"
	"github.com/codecrafters-io/http-server-starter-go/app/pkg/routing"
)

var _ = net.Listen
var _ = os.Exit
var REQ_END = "\r\n"

func root(conn net.Conn, request *models.Request) {
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n\r\n")
}

func (s *Server) echo(conn net.Conn, request *models.Request) {
	parts := request.UrlParts

	if len(parts) < 2 {
		fmt.Fprint(conn, "HTTP/1.1 400 Bad Request\r\n\r\n")
		return
	}

	response := fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s", len(parts[2]), parts[2])
	fmt.Fprint(conn, response)
	return

	// fmt.Fprint(conn, "HTTP/1.1 404 Not Found\r\n\r\n")
}

func main() {
	// Setup Routers
	router := routing.NewRouterMap()
	router.AssignHandler("/", root)

	// Setup Server
	server := Server{
		router: router,
	}

	router.AssignHandler("/echo", server.echo)
	server.Start()
}

type Server struct {
	listener net.Listener
	router   *routing.Router
}

func (s *Server) Start() {
	s.Listen()
	defer s.Close()
	for {
		conn := s.Accept()
		go s.handlerConnection(conn)
	}

}

func (s *Server) Listen() {
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	s.listener = l
}

func (s *Server) Accept() net.Conn {
	conn, err := s.listener.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}
	return conn
}
func (s *Server) Close() {
	err := s.listener.Close()
	if err != nil {
		fmt.Println("Failed to close listener:", err.Error())
	}

}

// TODO: make a context struct that takes method, path, headers, body.
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

		if strings.Contains(builder.String(), "\r\n\r\n") {
			break
		}
	}

	requestString := builder.String()
	return requestString
}

// and then make all the request funcs use it
func (s *Server) Read(conn net.Conn) models.Request {
	conn_string := connectionToString(conn)
	iter := strings.SplitSeq(conn_string, REQ_END)

	var request models.Request
	for partString := range iter {
		parts := strings.Split(partString, " ")

		switch strings.ToLower(parts[0]) {
		case "get":

			request.Method = parts[0]
			request.URL = parts[1]
			request.Body = parts[2]
		default:
			continue
			// fmt.Println("hi")
		}
	}
	request.UrlParts = strings.Split(request.URL, "/")
	return request

}

func (s *Server) handlerConnection(conn net.Conn) {
	defer conn.Close()
	read := s.Read(conn)
	s.router.Serve(conn, &read)
}
