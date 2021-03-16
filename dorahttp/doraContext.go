package dorahttp

import (
	"strconv"
	"time"
)

// Dctx represents the Context which hold the HTTP request and response.
// It has methods for the request query string, parameters, body, HTTP headers and so on.
type Dctx struct {
	Protocol string            // HTTP type
	Method   string            // HTTP method
	Path     string            // HTTP path
	Params   map[string]string // All  parameters
	Query    string            // Query string
	Head     string
	Body     string
	IP       string
}

// Cookie data for c.Cookie
type Cookie struct {
	Name     string    `json:"name"`
	Value    string    `json:"value"`
	Path     string    `json:"path"`
	Domain   string    `json:"domain"`
	MaxAge   int       `json:"max_age"`
	Expires  time.Time `json:"expires"`
	Secure   bool      `json:"secure"`
	HTTPOnly bool      `json:"http_only"`
	SameSite string    `json:"same_site"`
}

//DoraHandler ==> func for handle a request
//return a Dtcx.Response
type DoraHandler func(ctx *Dctx) []byte

//Response is  http_response for server
func (dctx *Dctx) Response(statusCode int, res string) []byte {
	return appendResp(string(statusCode)+" "+StatusText(statusCode), "", res)
}

func selfCallResponse(statusCode int, res string) []byte {
	return appendResp(string(statusCode)+" "+StatusText(statusCode), "", res)
}

// appendResp will append a valid http response to the provide bytes.
// The head parameter should be a series of lines ending with "\r\n" or empty.

func appendResp(status, head, body string) []byte {
	var b []byte
	b = append(b, "HTTP/1.1"...)
	b = append(b, ' ')
	b = append(b, status...)
	b = append(b, '\r', '\n')
	b = append(b, "Server: Dora\r\n"...)
	b = append(b, "Date: "...)
	b = time.Now().AppendFormat(b, "Mon, 02 Jan 2006 15:04:05 GMT")
	b = append(b, '\r', '\n')
	if len(body) > 0 {
		b = append(b, "Content-Length: "...)
		b = strconv.AppendInt(b, int64(len(body)), 10)
		b = append(b, '\r', '\n')
	}
	b = append(b, head...)
	b = append(b, '\r', '\n')
	if len(body) > 0 {
		b = append(b, body...)
	}
	return b
}
