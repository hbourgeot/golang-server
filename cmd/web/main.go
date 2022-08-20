package main

import "log"

func isAnError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	server := NewServer(":4000")
	server.Handle("GET", "/", HandleRoot)
	server.Handle("POST", "/api", server.AddMiddleware(HandleHome, CheckAuth(), Loggin()))
	server.Handle("POST", "/create", PostRequest)
	server.Handle("POST", "/user", UserPostRequest)

	err := server.Listen()
	isAnError(err)

}
