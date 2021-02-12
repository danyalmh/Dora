package dorahttp

import (
	"strconv"

	"dora.com/fullstack/framework/allocator"
)

//HTTPRequest is context of Request
type HTTPRequest struct {
	Method      string
	URI         string
	HTTPVersion string
	Token       string
}

//ResponseB is Response with Parameters of []byte type
func ResponseB(statusCode []byte, statusText []byte, contentType []byte,
	contentLength []byte, body []byte) []byte {

	return allocator.ConcatCopyPreAllocate([][]byte{
		[]byte("HTTP/1.1 "), statusCode, []byte(" "), statusText,
		[]byte("\r\nContent-Type: "), contentType,
		[]byte("\r\nContent-Length: "), contentLength,
		[]byte("\n\n"),
		body})

}

// I Think this have a error
// Because don't have (body & content-length)

//ResponseBNOBody is Response with Parameters of []byte type
// without have ( body )
func ResponseBNOBody(statusCode []byte, statusText []byte) []byte {

	return allocator.ConcatCopyPreAllocate([][]byte{
		[]byte("HTTP/1.1 "), statusCode, []byte(" "), statusText,
		[]byte("\r\nContent-Type: "), []byte("plain/text\n\n")})

}

//ResponseS Response with Parameters of string type
// FOR SIMPLER (((( but i don't RECOMMENDED IT ))))
func ResponseS(statusCode string, statusText string, contentType string, body string) []byte {

	return ResponseB([]byte(statusCode), []byte(statusText), []byte(contentType),
		[]byte(strconv.Itoa(len(body))), []byte(body))
}

//ResponseSNOBody Response with without body
// FOR SIMPLER (((( but i don't RECOMMENDED IT ))))
func ResponseSNOBody(statusCode string, statusText string) []byte {

	return ResponseBNOBody([]byte(statusCode), []byte(statusText))
}
