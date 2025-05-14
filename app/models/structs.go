package models

type Request struct {
	Method   string
	URL      string
	UrlParts []string
	Body     string
	Headers  map[string]string
}
