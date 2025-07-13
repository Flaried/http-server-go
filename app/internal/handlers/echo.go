package handlers

import (
	"fmt"
	"github.com/codecrafters-io/http-server-starter-go/app/internal/models"
	"net"
)

// Endpoints that returns the Query Params. Returns bad request to the client if theres no param.
func Echo(conn net.Conn, req models.Request) {
	fmt.Println("echo")
	// Extract path parameter
	param := models.QueryParam(req)
	if param == "" {
		resp := models.BadRequest()
		fmt.Fprint(conn, resp.String())
		return
	}

	echoText := req.Path[2]
	resp := models.OkStatus(&req, []byte(echoText), "text/plain")

	fmt.Fprint(conn, resp.String())
}
