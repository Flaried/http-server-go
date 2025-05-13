package routing

import (
	"fmt"
	"net"
	"strings"
)

type Context {
	
}
// The string is the request string
type HandlerFunc func(net.Conn)

// Router is a dictionary of strings and handler functions
type Router struct {
	routes map[string]HandlerFunc
}

func (r *Router) AssignHandler(path string, handler HandlerFunc) {
	r.routes[path] = handler
}

func (r *Router) Serve(conn net.Conn, dataRecieved string) {
	fullPath := strings.Split(dataRecieved, " ")[1]
	for routePath, routeFunc := range r.routes {
		if fullPath == routePath || strings.HasPrefix(fullPath, routePath+"/") {
			routeFunc(conn, &dataRecieved)
			return
		}
	}
	fmt.Fprintf(conn, "HTTP/1.1 404 Not Found\r\n\r\n")
}

func NewRouterMap() *Router {
	return &Router{
		routes: make(map[string]HandlerFunc),
	}
}
