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
	router.AssignHandler("/user-agent", network.Root)

	// Setup Server
	server := network.Server{
		Router: router,
	}

	router.AssignHandler("/echo", server.Echo)
	server.Start()
}
