package routing

import (
	"fmt"
	"net"
	"strings"

	"github.com/codecrafters-io/http-server-starter-go/app/models"
)

// The string is the request string
type HandlerFunc func(net.Conn, *models.Request)

// Router is a dictionary of strings and handler functions
type Router struct {
	routes map[string]HandlerFunc
}

func (r *Router) AssignHandler(path string, handler HandlerFunc) {
	r.routes[path] = handler
}

func (r *Router) Serve(conn net.Conn, request *models.Request) {
	for routePath, routeFunc := range r.routes {
		fmt.Println(request.UrlParts[1], "Route", routePath)
		if request.UrlParts[1] == strings.Replace(routePath, "/", "", 1) {
			routeFunc(conn, request)
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
