package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

var _ = net.Listen
var _ = os.Exit

func main() {
	s := Server{}
	s.Start()
}

type Server struct {
	listener net.Listener
}

func (s *Server) Start() {
	s.Listen()
	defer s.Close()
	for {
		conn := s.Accept()
		go handlerConnection(conn)
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

func handlerConnection(conn net.Conn) {
	defer conn.Close()
	resp, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Printf("Error reading connect: %v", err)
		os.Exit(1)
	}

	requestPath := strings.Split(resp, " ")[1]

	if requestPath == "/" {
		fmt.Fprintf(conn, "HTTP/1.1 200 OK\r\n\r\n")
	} else {
		fmt.Fprintf(conn, "HTTP/1.1 404 Not Found\r\n\r\n")
	}
	// response := "HTTP/1.1 200 OK\r\n\r\n"
	// conn.Write([]byte(response))
}
