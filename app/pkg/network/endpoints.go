package network

import (
	"fmt"
	"net"

	"github.com/codecrafters-io/http-server-starter-go/app/models"
)

func Root(conn net.Conn, request *models.Request) {
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n\r\n")
}

func (s *Server) Echo(conn net.Conn, request *models.Request) {
	parts := request.UrlParts
	if len(parts) < 2 {
		fmt.Fprint(conn, "HTTP/1.1 400 Bad Request\r\n\r\n")
		return
	}

	response := fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s", len(parts[2]), parts[2])
	fmt.Fprint(conn, response)
	return

}

func (s *Server) UserAgent(conn net.Conn, request *models.Request) {
	agent := request.Headers["User-Agent"]

	response := fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s", len(agent), agent)
	fmt.Fprint(conn, response)
}
