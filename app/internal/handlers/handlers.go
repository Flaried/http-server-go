package handlers

import (
	"github.com/codecrafters-io/http-server-starter-go/app/internal/models"
	"net"
)

// Handler interface - now owned by handlers package
type Handler interface {
	Handle(conn net.Conn, req models.Request)
}

// HandlerFunc type for function handlers
type HandlerFunc func(conn net.Conn, req models.Request)

func (f HandlerFunc) Handle(conn net.Conn, req models.Request) {
	f(conn, req)
}
