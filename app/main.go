package main

import (
	"fmt"
	"net"
	"os"

	"github.com/codecrafters-io/http-server-starter-go/app/internal/constants"
	"github.com/codecrafters-io/http-server-starter-go/app/internal/server"
)

var _ = net.Listen
var _ = os.Exit

func main() {
	// Setup Routers
	router := server.NewRouterMap()
	router.AssignHandler("/", server.Root)

	// Setup Server
	server := server.Server{
		Router: router,
	}

	router.AssignHandler("/echo", server.Echo)
	router.AssignHandler("/user-agent", server.UserAgent)
	config := constants.ServerConfig{
		Address:  "0.0.0.0:4221",
		Protocol: "tcp",
	}

	fmt.Printf("Server Listening %s\n", config.Address)
	server.Start(config)
}
