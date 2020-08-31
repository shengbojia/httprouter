package httprouter

import (
	"fmt"
	"net/http"
)

// Implements http.Handler interface, so it can serve http requests
type Router struct {
	handlers map[string]func(http.ResponseWriter, *http.Request)
}

// returns a new instance of Router
func NewRouter() *Router {
	return &Router{handlers: make(map[string]func(http.ResponseWriter, *http.Request))}
}

func (r *Router) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	f, found := r.handlers[toKey(request.Method, request.URL.Path)]
	if !found {
		notFound(writer)
		return
	}

	f(writer, request)
}

func (r *Router) GET(path string, f http.HandlerFunc) {
	r.handlers[toKey("GET", path)] = f
}

func (r *Router) POST(path string, f http.HandlerFunc) {
	r.handlers[toKey("POST", path)] = f
}

func (r *Router) DELETE(path string, f http.HandlerFunc) {
	r.handlers[toKey("DELETE", path)] = f
}

func (r *Router) PUT(path string, f http.HandlerFunc) {
	r.handlers[toKey("PUT", path)] = f
}

func toKey(method, urlPath string) string {
	return fmt.Sprintf("%s:%s", method, urlPath)
}

func notFound(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
	writer.Write([]byte(`{"error":"not found"}`))
}







