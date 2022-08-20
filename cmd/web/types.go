package main

import "net/http"

type Middleware func(handlerFunc http.HandlerFunc) http.HandlerFunc
