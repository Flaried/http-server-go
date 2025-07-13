package handlers

import (
	"fmt"
	"github.com/codecrafters-io/http-server-starter-go/app/internal/models"
	"net"
)

func Echo(conn net.Conn, req models.Request) {
	// Extract path parameter
	param := models.QueryParam(req)
	if param != "" {
		resp := models.Response{
			StatusCode: 400,
			StatusText: "Bad Request",
			Headers:    map[string]string{},
			Body:       "",
		}
		fmt.Fprint(conn, resp.String())
		return
	}

	echoText := req.Path[2]
	resp := models.Response{
		StatusCode: 200,
		StatusText: "OK",
		Headers: map[string]string{
			"Content-Type": "text/plain",
		},
		Body: echoText,
	}
	fmt.Fprint(conn, resp.String())
}
