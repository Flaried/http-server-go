package routing

import (
	"fmt"
	"net"
)

type HandlerFunc func(net.Conn)

// Router is a dictionary of strings and handler functions
type Router struct {
	routes map[string]HandlerFunc
}

func (r *Router) AssignHandler(path string, handler HandlerFunc) {
	r.routes[path] = handler
}

func (r *Router) Serve(conn net.Conn, path string) {
	if handler, ok := r.routes[path]; ok {
		handler(conn)
	} else {
		fmt.Fprintf(conn, "HTTP/1.1 404 Not Found\r\n\r\n")
	}
}

func NewRouterMap() *Router {
	return &Router{
		routes: make(map[string]HandlerFunc),
	}
}
