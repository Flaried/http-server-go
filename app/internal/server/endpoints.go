package server

import (
	"fmt"
	"net"

	"github.com/codecrafters-io/http-server-starter-go/app/internal/config"
)

func Root(conn net.Conn, request *config.Request) {
	fmt.Fprint(conn, config.HTTPVersion+" 200 OK"+config.MARKER)
}

func (s *Server) Echo(conn net.Conn, request *config.Request) {
	parts := request.UrlParts
	if len(parts) < 3 {
		fmt.Fprint(conn, config.HTTPVersion+" 400 Bad Request"+config.MARKER)
		return
	}

	body := parts[2]
	response := fmt.Sprintf(
		"%s 200 OK%sContent-Type: text/plain%sContent-Length: %d%s%s%s",
		config.HTTPVersion, config.CRLF,
		config.CRLF,
		len(body), config.MARKER,
		body,
		config.CRLF, // optional final CRLF
	)
	fmt.Fprint(conn, response)
}

func (s *Server) UserAgent(conn net.Conn, request *config.Request) {
	agent := request.Headers["User-Agent"]

	response := fmt.Sprintf(
		"%s 200 OK%sContent-Type: text/plain%sContent-Length: %d%s%s%s",
		config.HTTPVersion, config.CRLF,
		config.CRLF,
		len(agent), config.MARKER,
		agent,
		config.CRLF, // optional final CRLF
	)
	fmt.Fprint(conn, response)
}
