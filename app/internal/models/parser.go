package models

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"errors"
	"io"
	"net"
	"strconv"
	"strings"
)

/*
ParseRequest reads and parses an HTTP request from the given connection.
It returns a Request struct with method, URL, path segments, headers, and optional body.
An error is returned if the request line is invalid or if reading the body fails.
*/
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

// QueryParam extracts the third segment of the request path, if present.
// Returns an empty string if the path has fewer than 3 segments.
func QueryParam(request Request) string {
	path := request.Path
	if len(path) < 3 {
		return ""
	}
	return path[2]
}

// Gzip compresses the given byte slice using gzip and returns the compressed data.
// If compression fails, an empty byte slice is returned.
func Gzip(body []byte) []byte {
	var b bytes.Buffer
	writer := gzip.NewWriter(&b)
	defer writer.Close()

	_, err := writer.Write(body)
	if err != nil {
		return []byte("")
	}

	writer.Close()
	return b.Bytes()
}
