package server

import (
	"fmt"
	"net"

	"github.com/codecrafters-io/http-server-starter-go/app/internal/constants"
)

func Root(conn net.Conn, request *constants.Request) {
	resp := constants.Response{
		StatusCode: 200,
		StatusText: "OK",
		Headers:    map[string]string{},
		Body:       "",
	}
	_, _ = fmt.Fprint(conn, resp.String())
}

func (s *Server) Echo(conn net.Conn, request *constants.Request) {
	parts := request.UrlParts
	if len(parts) < 3 {
		resp := constants.Response{
			StatusCode: 400,
			StatusText: "Bad Request",
			Headers:    map[string]string{},
			Body:       "",
		}
		_, _ = fmt.Fprint(conn, resp.String())
		return
	}

	resp := constants.Response{
		StatusCode: 200,
		StatusText: "OK",
		Headers: map[string]string{
			"Content-Type": "text/plain",
		},
		Body: parts[2],
	}
	_, _ = fmt.Fprint(conn, resp.String())
}

func (s *Server) UserAgent(conn net.Conn, request *constants.Request) {
	agent := request.Headers["User-Agent"]

	resp := constants.Response{
		StatusCode: 200,
		StatusText: "OK",
		Headers: map[string]string{
			"Content-Type": "text/plain",
		},
		Body: agent,
	}
	_, _ = fmt.Fprint(conn, resp.String())
}
