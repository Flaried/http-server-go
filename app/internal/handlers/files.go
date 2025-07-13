package handlers

import (
	"fmt"
	"github.com/codecrafters-io/http-server-starter-go/app/internal/models"
	"net"
	"os"
)

// The type to determine where to get the files
type FileHandler struct {
	ServingDirectory string
}

// Makes a new file handler with the directory to fetch from
func NewFileHandler(servingDir string) *FileHandler {
	return &FileHandler{ServingDirectory: servingDir}
}

// When the file API is GET the file to get is the QueryParam
func (h FileHandler) HandleGet(conn net.Conn, req models.Request) {
	if len(req.Path) < 3 {
		resp := models.BadRequest()
		fmt.Fprint(conn, resp.String())
		return
	}

	var filename string
	param := models.QueryParam(req)
	if param != "" {
		filename = param
	} else {
		resp := models.NotFound()
		fmt.Fprint(conn, resp.String())
		return
	}

	filePath := fmt.Sprintf("%s%s", h.ServingDirectory, filename)
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		resp := models.NotFound()
		fmt.Fprint(conn, resp.String())
		return
	}

	resp := models.OkStatus(&req, bytes, "application/octet-stream")
	fmt.Fprint(conn, resp.String())
}

// When the file API is POST the filename is the queryparam and the file data is the request data
func (h FileHandler) HandlePost(conn net.Conn, req models.Request) {
	if len(req.Body) == 0 {
		resp := models.BadRequest()
		fmt.Fprint(conn, resp.String())
		return
	}

	fmt.Println(string(req.Body))
	var filename string
	param := models.QueryParam(req)
	if param != "" {
		filename = param
	} else {
		resp := models.NotFound()
		fmt.Fprint(conn, resp.String())
		return
	}

	filePath := fmt.Sprintf("%s%s", h.ServingDirectory, filename)
	fmt.Println(filePath, "hee")
	err := os.WriteFile(filePath, req.Body, 0644)
	if err != nil {
		resp := models.InternalServerError()
		fmt.Fprint(conn, resp.String())
		return
	} else {
		fmt.Printf("Saved file in %s\n", filePath)
	}

	resp := models.Response{
		StatusCode: 201,
		StatusText: "Created",
		Headers:    map[string]string{},
		Body:       nil,
	}
	fmt.Fprint(conn, resp.String())
}
