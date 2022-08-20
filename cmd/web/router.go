package main

import (
	"net/http"
)

type Router struct {
	rules map[string]http.HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		rules: make(map[string]http.HandlerFunc),
	}
}

func (rou *Router) FindHandler(path string) (http.HandlerFunc, bool) {
	handler, exist := rou.rules[path]
	return handler, exist
}

func (rou *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler, exist := rou.FindHandler(r.URL.Path)

	if !exist {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	handler(w, r)
}
