package main

import (
	"bufio"
	"fmt"
	"github.com/codecrafters-io/http-server-starter-go/app/pkg/routing"
	"net"
	"os"
	"strings"
)

var _ = net.Listen
var _ = os.Exit

func root(conn net.Conn) {
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n\r\n")
}
func (s *Server) echo(conn net.Conn) {
	fmt.Println("reading")
	resp := s.Read(conn)
	fmt.Println("hae")
	parts := strings.Split(resp, " ")
	if len(parts) < 2 {
		fmt.Fprint(conn, "HTTP/1.1 400 Bad Request\r\n\r\n")
		return
	}

	path := parts[1]
	prefix := "/echo/"

	if strings.HasPrefix(path, prefix) {
		echoParam := strings.TrimPrefix(path, prefix)

		response := fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length :%d\r\n\r\n%s", len(echoParam), echoParam)
		fmt.Fprint(conn, response)
		fmt.Println("Sent back")
		return
	}

	fmt.Fprint(conn, "HTTP/1.1 404 Not Found\r\n\r\n")
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
//
//	and then make all the request funcs use it
func (s *Server) Read(conn net.Conn) string {
	fmt.Println("0")
	reader := bufio.NewReader(conn)
	var request string

	fmt.Println("1")
	for {
		fmt.Println("2")

		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error reading: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("3")
		fmt.Printf("Line: %q\n", line) // Debug output
		request += line

		fmt.Println("4")
		if line == "\r\n" || line == "\n" {
			fmt.Println("break")
			break // End of headers
		}
	}

	return request
}

func (s *Server) handlerConnection(conn net.Conn) {
	defer conn.Close()
	read := s.Read(conn)
	s.router.Serve(conn, read)
}
