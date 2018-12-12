package routes

import (
	"log"
	"net/http"
	"time"
)

type endpoint struct {
	path     string
	handlers map[string]http.HandlerFunc
}

// Router maintains configuration for an implementation of http.Handler.
type Router struct {
	endpoints map[string]*endpoint

	NotFoundHandler   http.HandlerFunc
	NotAllowedHandler http.HandlerFunc
}

// NewRouter returns a new initialized Router.
func NewRouter() *Router {
	return new(Router)
}

// Get adds a new handler for the GET HTTP method.
func (r *Router) Get(path string, handler http.HandlerFunc) {
	r.addHandler(http.MethodGet, path, handler)
}

// Post adds a new handler for the POST HTTP method.
func (r *Router) Post(path string, handler http.HandlerFunc) {
	r.addHandler(http.MethodPost, path, handler)
}

// Put adds a new handler for the Put HTTP method.
func (r *Router) Put(path string, handler http.HandlerFunc) {
	r.addHandler(http.MethodPut, path, handler)
}

// Delete adds a new handler for the Delete HTTP method.
func (r *Router) Delete(path string, handler http.HandlerFunc) {
	r.addHandler(http.MethodDelete, path, handler)
}

func (r *Router) addHandler(method string, path string, handler http.HandlerFunc) {
	if path[0] != '/' {
		panic("path must begin with '/' in path '" + path + "'")
	}

	if r.endpoints == nil {
		r.endpoints = make(map[string]*endpoint)
	}

	var p *endpoint
	p = r.endpoints[path]

	if p == nil {
		p = new(endpoint)
	}

	if p.handlers == nil {
		p.handlers = make(map[string]http.HandlerFunc)
	}

	p.path = path
	p.handlers[method] = handler

	r.endpoints[path] = p
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	method := req.Method
	endpoint := r.endpoints[path]

	start := time.Now()
	defer func() { log.Println(method, path, time.Since(start)) }()

	if endpoint == nil {
		if r.NotFoundHandler != nil {
			r.NotFoundHandler(w, req)
		} else {
			code := http.StatusNotFound
			http.Error(w, http.StatusText(code), code)
		}
		return
	}

	handler := endpoint.handlers[method]

	if handler == nil {
		if r.NotAllowedHandler != nil {
			r.NotAllowedHandler(w, req)
		} else {
			code := http.StatusMethodNotAllowed
			http.Error(w, http.StatusText(code), code)
		}
		return
	}

	handler(w, req)
}
