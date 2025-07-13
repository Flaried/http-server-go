package handlers

import (
	"github.com/codecrafters-io/http-server-starter-go/app/internal/models"
	"net"
)

type Handler interface {
	Handle(conn net.Conn, req models.Request)
}

// HandlerFunc type for function handlers
type HandlerFunc func(conn net.Conn, req models.Request)

// run the endpoint function with the connection and request.
func (f HandlerFunc) Handle(conn net.Conn, req models.Request) {
	f(conn, req)
}
