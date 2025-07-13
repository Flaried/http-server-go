package handlers

import (
	"fmt"
	"github.com/codecrafters-io/http-server-starter-go/app/internal/models"
	"net"
)

func UserAgent(conn net.Conn, req models.Request) {
	agent := req.Headers["user-agent"]

	resp := models.OkStatus(&req, []byte(agent), "text/plain")
	fmt.Fprint(conn, resp.String())
}
