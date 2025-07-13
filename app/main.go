package main

import (
	"flag"
	"fmt"
	"github.com/codecrafters-io/http-server-starter-go/app/internal/handlers"
	"github.com/codecrafters-io/http-server-starter-go/app/internal/models"
	"github.com/codecrafters-io/http-server-starter-go/app/internal/router"
	"github.com/codecrafters-io/http-server-starter-go/app/internal/server"
	"net"
	"os"
)

var path *string

func main() {
	config := models.ServerConfig{
		Address:  "0.0.0.0:4221",
		Protocol: "tcp",
	}

	server := server.NewServer(config)

	r := router.NewRouter()

	path = flag.String("directory", ".", "Directory to serve files from")
	flag.Parse()
	fileHandler := handlers.NewFileHandler(*path)

	// Root Path
	r.RegisterFunc("GET", "/", func(conn net.Conn, req models.Request) {
		resp := models.Response{
			StatusCode: 200,
			StatusText: "OK",
			Headers:    map[string]string{},
			Body:       nil,
		}
		fmt.Fprint(conn, resp.String())
	})

	// Routes
	r.RegisterFunc("GET", "/echo", handlers.Echo)
	r.RegisterFunc("GET", "/user-agent", handlers.UserAgent)
	r.RegisterFunc("GET", "/files", handlers.HandlerFunc(fileHandler.HandleGet))
	r.RegisterFunc("POST", "/files", handlers.HandlerFunc(fileHandler.HandlePost))

	// Set router and start server
	server.SetRouter(r)
	fmt.Println("Server starting on", config.Address)
	if err := server.Start(); err != nil {
		fmt.Printf("Server failed to start: %v\n", err)
		os.Exit(1)
	}
}
