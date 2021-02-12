package dora

import (
	"fmt"
	"net"
	"strings"

	"dora.com/fullstack/framework/dorahttp"
	"dora.com/fullstack/framework/router"
)

var (
	rtx *router.Router
)

//Start initialize all component for running ....
func Start(port string, rt *router.Router) {

	rtx = rt

	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		// handle error
		panic("cannot open port: " + port)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}
		go handleConnection(conn)

	}
}

func handleConnection(conn net.Conn) {
	buff := make([]byte, 1024)

	_, err := conn.Read(buff)
	if err != nil {
		return
	}

	reqx, statusCode, statusText := parse(buff)
	if statusCode != nil {
		conn.Write(dorahttp.ResponseBNOBody(statusCode, statusText))
		return
	}

	handleRequest, err := rtx.Extract(reqx.Method, reqx.URI)
	fmt.Println(string(dorahttp.ResponseBNOBody([]byte("405"), []byte("Method Not Allowed"))))
	if err != nil {
		conn.Write(dorahttp.ResponseBNOBody([]byte("405"), []byte("Method Not Allowed")))
		return
	}
	responseByte := handleRequest(reqx)
	conn.Write(responseByte)
}

func parse(data []byte) (*dorahttp.HTTPRequest, []byte, []byte) {

	req := new(dorahttp.HTTPRequest)
	datas := string(data)
	lines := strings.Split(datas, "/r/n")
	requestLine := lines[0]
	words := strings.Split(requestLine, " ")

	switch words[0] {
	case "DELETE":
		req.Method = "DELETE"
	case "GET":
		req.Method = "GET"
	case "POST":
		req.Method = "POST"
	case "PUT":
		req.Method = "PUT"
	case "PATCH":
		req.Method = "PATCH"
	default:
		return req, []byte("400"), []byte("Bad Request")
	}

	if len(words[1]) <= 1 {
		return req, []byte("400"), []byte("Bad Request")
	}
	req.URI = words[1]

	req.HTTPVersion = words[2] + " " + words[3]
	// Request Header is Ok

	return req, nil, nil
}
