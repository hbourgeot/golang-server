package main

import (
	"net/http"
)

type Router struct {
	rules map[string]map[string]http.HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		rules: make(map[string]map[string]http.HandlerFunc),
	}
}

func (rou *Router) FindHandler(path, method string) (http.HandlerFunc, bool, bool) {
	_, exist := rou.rules[path]
	handler, methodExists := rou.rules[path][method]
	return handler, methodExists, exist
}

func (rou *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler, methodExist, exist := rou.FindHandler(r.URL.Path, r.Method)

	if !exist {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if !methodExist {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	handler(w, r)
}
