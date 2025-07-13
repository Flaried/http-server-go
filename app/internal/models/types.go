package models

import (
	"fmt"
	"strings"
)

type Request struct {
	Method  string
	URL     string
	Path    []string
	Body    []byte
	Headers map[string]string
}

type Response struct {
	StatusCode int
	StatusText string
	Headers    map[string]string
	Body       []byte
	Params     *string
}

func InternalServerError() Response {
	return Response{
		StatusCode: 500,
		StatusText: "Internal Server Error",
		Headers:    map[string]string{},
		Body:       nil,
	}
}

func OkStatus(req *Request, body []byte, contentType string) Response {
	respHeaders := make(map[string]string)

	acceptEndcoding := strings.ToLower(req.Headers["accept-encoding"])
	if strings.Contains(acceptEndcoding, "gzip") {
		body = Gzip(body)
		respHeaders["content-encoding"] = "gzip"
	}

	respHeaders["Content-Type"] = contentType

	return Response{
		StatusCode: 200,
		StatusText: "OK",
		Headers:    respHeaders,
		Body:       body,
	}
}

func BadRequest() Response {
	return Response{
		StatusCode: 400,
		StatusText: "Bad Request",
		Headers:    map[string]string{},
		Body:       nil,
	}
}
func NotFound() Response {
	return Response{
		StatusCode: 404,
		StatusText: "Not Found",
		Headers:    map[string]string{},
		Body:       nil,
	}
}

func (r Response) String() string {
	statusLine := fmt.Sprintf("%s %d %s%s", HTTPVersion, r.StatusCode, r.StatusText, CRLF)

	headers := ""
	for k, v := range r.Headers {
		headers += fmt.Sprintf("%s: %s%s", k, v, CRLF)
	}

	// Add Content-Length header automatically if not present
	if _, exists := r.Headers["Content-Length"]; !exists {
		headers += fmt.Sprintf("Content-Length: %d%s", len(r.Body), CRLF)
	}

	return statusLine + headers + CRLF + string(r.Body)
}
