package models

import "fmt"

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
	Body       string
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

	return statusLine + headers + CRLF + r.Body
}
