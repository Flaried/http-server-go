package models

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"
)

func ParseRequest(conn net.Conn) (Request, error) {
	var req Request
	reader := bufio.NewReader(conn)

	// Parse request line
	requestLine, err := reader.ReadString('\n')
	if err != nil {
		return req, err
	}

	parts := strings.Fields(requestLine)
	if len(parts) < 3 {
		return req, errors.New("invalid request line")
	}

	req.Method = parts[0]
	req.URL = parts[1]
	req.Path = strings.Split(parts[1], "/")

	// Parse headers
	headers := make(map[string]string)
	for {
		line, err := reader.ReadString('\n')
		if err != nil || line == "\r\n" {
			break
		}

		headerParts := strings.SplitN(line, ":", 2)
		if len(headerParts) == 2 {
			key := strings.TrimSpace(strings.ToLower(headerParts[0]))
			value := strings.TrimSpace(headerParts[1])
			headers[key] = value
		}
	}

	req.Headers = headers

	// Parse body if present
	if lengthStr := headers["content-length"]; lengthStr != "" {
		contentLength, err := strconv.Atoi(lengthStr)
		if err != nil {
			return req, err
		}

		body := make([]byte, contentLength)
		_, err = io.ReadFull(reader, body)
		if err != nil {
			return req, err
		}

		req.Body = body
	}
	return req, nil
}

func QueryParam(request Request) string {
	path := request.Path
	if len(path) < 3 {
		return ""
	}
	return path[2]
}
func Gzip(body []byte) []byte {
	fmt.Println("in gzip")
	var b bytes.Buffer
	writer := gzip.NewWriter(&b)
	defer writer.Close()

	_, err := writer.Write([]byte(body))
	if err != nil {
		return []byte("")
	}

	return b.Bytes()
}
