package models

// Default Configs and HTTP 1.1 consts
const (
	DefaultPort = "4221"
	DefaultHost = "0.0.0.0"
	HTTPVersion = "HTTP/1.1"
	CRLF        = "\r\n"
	MARKER      = CRLF + CRLF
)

// ServerConfig holds the configuration for the server such as address and protocol.
type ServerConfig struct {
	Address  string `default:"0.0.0.0:4221"`
	Protocol string `default:"tcp"`
}
