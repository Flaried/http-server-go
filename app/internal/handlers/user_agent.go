package handlers

import (
	"fmt"
	"github.com/codecrafters-io/http-server-starter-go/app/internal/models"
	"net"
)

func UserAgent(conn net.Conn, req models.Request) {
	agent := req.Headers["user-agent"]

	resp := models.Response{
		StatusCode: 200,
		StatusText: "OK",
		Headers: map[string]string{
			"Content-Type": "text/plain",
		},
		Body: agent,
	}
	fmt.Fprint(conn, resp.String())
}
