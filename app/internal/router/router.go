package router

import (
	"fmt"
	"github.com/codecrafters-io/http-server-starter-go/app/internal/handlers"
	"github.com/codecrafters-io/http-server-starter-go/app/internal/models"
	"net"
	"strings"
)

type RouteKey struct {
	Method string
	Path   string
}

type Router struct {
	routes map[RouteKey]handlers.Handler
}

func NewRouter() *Router {
	return &Router{
		routes: make(map[RouteKey]handlers.Handler),
	}
}

// helper func to map the handlerfunc with the path and method
func (r *Router) register(method, path string, handler handlers.Handler) {
	key := RouteKey{Method: method, Path: path}
	r.routes[key] = handler
}

// RegisterFunc is a helper to register a HandlerFunc directly as a route.
func (r *Router) RegisterFunc(method, path string, handlerFunc handlers.HandlerFunc) {
	r.register(method, path, handlerFunc)
}

// Serve routes the given request to the appropriate handler or returns a 404 response.
func (r *Router) Serve(conn net.Conn, req models.Request) {
	for routeKey, handler := range r.routes {
		if r.matchRoute(routeKey, req) {
			handler.Handle(conn, req)
			return
		}
	}

	// 404 Not Found
	resp := models.NotFound()
	fmt.Fprint(conn, resp.String())
}

// helper func that checks whether the given request matches a registered route.
func (r *Router) matchRoute(routeKey RouteKey, req models.Request) bool {
	if req.Method != routeKey.Method {
		return false
	}

	if len(req.Path) > 1 {
		requestPath := strings.Replace(req.Path[1], "/", "", 1)
		routePath := strings.Replace(routeKey.Path, "/", "", 1)
		return requestPath == routePath
	}

	return routeKey.Path == "/"
}
