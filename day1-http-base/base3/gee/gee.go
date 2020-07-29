package gee

import (
	"fmt"
	"net/http"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)

type Engine struct {
	router map[string]HandlerFunc
}

func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

func (engine *Engine) addRoute(method string, pattren string, handler HandlerFunc) {
	key := method + "_" + pattren
	engine.router[key] = handler
}

func (engine *Engine) GET(pattren string, handler HandlerFunc) {
	engine.addRoute("GET", pattren, handler)
}

func (engine *Engine) POST(pattren string, handler HandlerFunc) {
	engine.addRoute("POST", pattren, handler)
}

func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "_" + req.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(w, req)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}
