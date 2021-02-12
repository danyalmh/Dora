package main

import (
	"strconv"

	"dora.com/fullstack/framework/dora"
	"dora.com/fullstack/framework/dorahttp"
	"dora.com/fullstack/framework/htmlhelper"
	"dora.com/fullstack/framework/router"
)

var tempTest = new(container)

func main() {

	rt := router.NewRouter()
	rt.Add("GET", "/dora/index", handleRequest2)
	rt.Add("GET", "/dora/getval", handleRequest)
	rt.Add("POST", "/dora/postval", handleRequest)

	dora.Start("8089", rt)
	/*
		Here Must Define All Function for
		setting to router for response

	*/

}

// ==================================>> (result of ResponseB)
func handleRequest(req *dorahttp.HTTPRequest) []byte {
	switch req.Method {
	case "GET":
		return dorahttp.ResponseS("200", "OK", "text/html", tempTest.Get())
	case "POST":
		tempTest.Set("New Value")
		return dorahttp.ResponseSNOBody("200", "OK")
	}

	return dorahttp.ResponseSNOBody("400", "Bad Request")
}

func handleRequest2(req *dorahttp.HTTPRequest) []byte {

	body := generateHTTPBody()
	return dorahttp.ResponseB([]byte("200"), []byte("OK"), []byte("html"), []byte(strconv.Itoa(len(body))), body)
}

/// ================================

func generateHTTPBody() []byte {
	return htmlhelper.HTMLGenerator("", htmlhelper.Div("Did", "Dname", "font-size:medium; margin:auto",
		"", htmlhelper.A("Aid", "Aname", "#", "margin:auto", "", "==== awesome Dora ====")+
			htmlhelper.Img("Iid", "Iname", "localhost:8080/quite.png", "width:500px;height:500px", "")))
}

type icontainer interface {
	Set(value string)
	Get() string
}

type container struct {
	value string
}

func (c *container) Set(value string) {
	c.value = value
}

func (c *container) Get() string {
	return c.value
}
