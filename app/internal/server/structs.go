package server

import (
	"fmt"
	"github.com/codecrafters-io/http-server-starter-go/app/internal/constants"
)

type Request struct {
	Method  string
	URL     string
	Path    []string
	Body    string
	Headers map[string]string
}

type Response struct {
	StatusCode int
	StatusText string
	Headers    map[string]string
	Body       string
}

func (r *Response) String() string {
	statusLine := fmt.Sprintf("%s %d %s%s", constants.HTTPVersion, r.StatusCode, r.StatusText, constants.CRLF)

	headers := ""
	for k, v := range r.Headers {
		headers += fmt.Sprintf("%s: %s%s", k, v, constants.CRLF)
	}

	// Add Content-Length header automatically if not present
	if _, exists := r.Headers["Content-Length"]; !exists {
		headers += fmt.Sprintf("Content-Length: %d%s", len(r.Body), constants.CRLF)
	}

	return statusLine + headers + constants.CRLF + r.Body
}

func (s *Server) GetQueryParam(request *Request) (string, *Response) {
	path := request.Path
	if len(path) < 3 {
		resp := Response{
			StatusCode: 400,
			StatusText: "Bad Request",
			Headers:    map[string]string{},
			Body:       "",
		}
		return "", &resp
	}
	return path[2], nil
}
