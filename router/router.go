package router

import (
	"errors"
	"fmt"

	"dora.com/fullstack/framework/dorahttp"
)

// ****** <IF HAPPEN DATA RACE USE mutex.RLOCK> *******

type iRouter interface {
	Add(method string, uri string, hand func(req *dorahttp.HTTPRequest) []byte)
	Extract(method string, uri string) func(req *dorahttp.HTTPRequest) []byte
	Print()
}

//HandleRequest handle a request and return (statusCode, statusText)

//Router is strcut for keppping all HTTPRequest
type Router struct {
	storage map[string]func(req *dorahttp.HTTPRequest) []byte
}

//NewRouter return a new router context
func NewRouter() *Router {
	return &Router{
		storage: make(map[string]func(req *dorahttp.HTTPRequest) []byte)}
}

//Add add (method, uri, handleRequest) to router
func (router *Router) Add(method string, uri string, hand func(req *dorahttp.HTTPRequest) []byte) {
	router.storage[method+uri] = hand

	fmt.Println("===>", router.storage)
}

//Extract if exist return handler
func (router *Router) Extract(method string, uri string) (func(req *dorahttp.HTTPRequest) []byte, error) {
	fmt.Println("===>", router.storage)
	val, ok := router.storage[method+uri]
	if ok != true {
		return nil, errors.New("Not Found")
	}

	return val, nil

}

func (r *Router) Print() {
	fmt.Println("===>", r.storage)
}
