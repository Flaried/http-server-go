package main

import (
	"github.com/codecrafters-io/http-server-starter-go/app/pkg/network"
	"github.com/codecrafters-io/http-server-starter-go/app/pkg/routing"
	"net"
	"os"
)

var _ = net.Listen
var _ = os.Exit

func main() {
	// Setup Routers
	router := routing.NewRouterMap()
	router.AssignHandler("/", network.Root)

	// Setup Server
	server := network.Server{
		Router: router,
	}

	router.AssignHandler("/echo", server.Echo)
	router.AssignHandler("/user-agent", server.UserAgent)
	server.Start()
}
