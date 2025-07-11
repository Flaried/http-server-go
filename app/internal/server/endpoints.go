package server

import (
	"fmt"
	"net"
	"os"
)

func Root(conn net.Conn, request *Request) {
	resp := Response{
		StatusCode: 200,
		StatusText: "OK",
		Headers:    map[string]string{},
		Body:       "",
	}
	_, _ = fmt.Fprint(conn, resp.String())
}

func (s *Server) Echo(conn net.Conn, request *Request) {
	queryParameter, badResponse := s.GetQueryParam(request)
	if badResponse != nil {
		fmt.Fprint(conn, badResponse.String())
		return
	}

	resp := Response{
		StatusCode: 200,
		StatusText: "OK",
		Headers: map[string]string{
			"Content-Type": "text/plain",
		},
		Body: queryParameter,
	}
	_, _ = fmt.Fprint(conn, resp.String())
}

func (s *Server) UserAgent(conn net.Conn, request *Request) {
	agent := request.Headers["User-Agent"]

	resp := Response{
		StatusCode: 200,
		StatusText: "OK",
		Headers: map[string]string{
			"Content-Type": "text/plain",
		},
		Body: agent,
	}
	_, _ = fmt.Fprint(conn, resp.String())
}

func (s *Server) ReturnFile(conn net.Conn, request *Request) {
	queryParameter, badResponse := s.GetQueryParam(request)
	if badResponse != nil {
		fmt.Fprint(conn, badResponse.String())
		return
	}

	file_name := queryParameter

	filePath := fmt.Sprintf("%s%s", *s.ServingDirectory, file_name)
	fmt.Println(filePath)
	bytes, err := os.ReadFile(filePath)

	if err != nil {
		resp := Response{
			StatusCode: 404,
			StatusText: "Not Found",
			Headers:    map[string]string{},
			Body:       "",
		}

		_, _ = fmt.Fprint(conn, resp.String())
		return
	}

	resp := Response{
		StatusCode: 200,
		StatusText: "OK",
		Headers: map[string]string{
			"Content-Type": "application/octet-stream",
		},
		Body: string(bytes),
	}

	_, _ = fmt.Fprint(conn, resp.String())
	// buf := make([]byte, 1024)

}
