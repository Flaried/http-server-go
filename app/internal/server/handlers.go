package server

import "net"
import "strings"
import "fmt"
import "github.com/codecrafters-io/http-server-starter-go/app/internal/constants"

// The string is the request string
type HandlerFunc func(net.Conn, *constants.Request)

// Router is a dictionary of strings and handler functions
type Router struct {
	routes map[string]HandlerFunc
}

func (r *Router) AssignHandler(path string, handler HandlerFunc) {
	r.routes[path] = handler
}

func (r *Router) Serve(conn net.Conn, request *constants.Request) {
	// fmt.Println(r.routes, "routes")
	for routePath, routeFunc := range r.routes {
		// fmt.Println(request.UrlParts[1], "Route", routePath)
		if request.UrlParts[1] == strings.Replace(routePath, "/", "", 1) {
			fmt.Println("Running function for", routePath)
			routeFunc(conn, request)
			return
		}
	}
	fmt.Fprintf(conn, "%s 404 Not Found%s", constants.HTTPVersion, constants.MARKER)

}

func NewRouterMap() *Router {
	return &Router{
		routes: make(map[string]HandlerFunc),
	}
}
