package models

const (
	DefaultPort = "4221"
	DefaultHost = "0.0.0.0"
	HTTPVersion = "HTTP/1.1"
	CRLF        = "\r\n"
	MARKER      = CRLF + CRLF
)

type ServerConfig struct {
	Address  string `default:"0.0.0.0:4221"`
	Protocol string `default:"tcp"`
}
