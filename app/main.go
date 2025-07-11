package main

import (
	"flag"
	"fmt"
	"github.com/codecrafters-io/http-server-starter-go/app/internal/constants"
	"github.com/codecrafters-io/http-server-starter-go/app/internal/server"
	"net"
	"os"
)

var _ = net.Listen
var _ = os.Exit
var path *string

func main() {
	path = flag.String("directory", ".", "Directory to serve files from")
	flag.Parse()

	fmt.Println("Serving files from directory:", *path)

	// Setup Routers

	router := server.NewRouterMap()
	router.AssignHandler("/", server.Root)

	// Setup Server
	server := server.Server{
		Router:           router,
		ServingDirectory: path,
	}

	router.AssignHandler("/echo", server.Echo)
	router.AssignHandler("/user-agent", server.UserAgent)
	config := constants.ServerConfig{
		Address:  "0.0.0.0:4221",
		Protocol: "tcp",
	}

	router.AssignHandler("/files", server.ReturnFile)
	fmt.Printf("Server Listening %s\n", config.Address)
	server.Start(config)
}
