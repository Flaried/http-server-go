package handlers

import (
	"fmt"
	"github.com/codecrafters-io/http-server-starter-go/app/internal/models"
	"net"
	"os"
)

type FileHandler struct {
	ServingDirectory string
}

func NewFileHandler(servingDir string) *FileHandler {
	return &FileHandler{ServingDirectory: servingDir}
}

func (h FileHandler) HandleGet(conn net.Conn, req models.Request) {
	if len(req.Path) < 3 {
		resp := models.Response{
			StatusCode: 400,
			StatusText: "Bad Request",
			Headers:    map[string]string{},
			Body:       "",
		}
		fmt.Fprint(conn, resp.String())
		return
	}

	var filename string
	param := models.QueryParam(req)
	if param != "" {
		filename = param
	} else {
		resp := models.Response{
			StatusCode: 404,
			StatusText: "Not Found",
			Headers:    map[string]string{},
			Body:       "",
		}
		fmt.Fprint(conn, resp.String())
		return
	}

	filePath := fmt.Sprintf("%s%s", h.ServingDirectory, filename)
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		resp := models.Response{
			StatusCode: 404,
			StatusText: "Not Found",
			Headers:    map[string]string{},
			Body:       "",
		}
		fmt.Fprint(conn, resp.String())
		return
	}

	resp := models.Response{
		StatusCode: 200,
		StatusText: "OK",
		Headers: map[string]string{
			"Content-Type": "application/octet-stream",
		},
		Body: string(bytes),
	}
	fmt.Fprint(conn, resp.String())
}

func (h FileHandler) HandlePost(conn net.Conn, req models.Request) {
	if len(req.Body) == 0 {
		resp := models.Response{
			StatusCode: 400,
			StatusText: "Bad Request",
			Headers:    map[string]string{},
			Body:       "",
		}
		fmt.Fprint(conn, resp.String())
		return
	}

	fmt.Println(string(req.Body))
	var filename string
	param := models.QueryParam(req)
	if param != "" {
		filename = param
	} else {
		resp := models.Response{
			StatusCode: 404,
			StatusText: "Not Found",
			Headers:    map[string]string{},
			Body:       "",
		}
		fmt.Fprint(conn, resp.String())
		return
	}

	filePath := fmt.Sprintf("%s%s", h.ServingDirectory, filename)
	fmt.Println(filePath, "hee")
	err := os.WriteFile(filePath, req.Body, 0644)
	if err != nil {
		resp := models.Response{
			StatusCode: 500,
			StatusText: "Internal Server Error",
			Headers:    map[string]string{},
			Body:       "",
		}
		fmt.Fprint(conn, resp.String())
		return
	} else {
		fmt.Printf("Saved file in %s\n", filePath)
	}

	resp := models.Response{
		StatusCode: 201,
		StatusText: "Created",
		Headers:    map[string]string{},
		Body:       "",
	}
	fmt.Fprint(conn, resp.String())
}
