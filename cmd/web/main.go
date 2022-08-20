package main

import "log"

func isAnError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	server := NewServer(":4000")
	server.Handle("/", HandleRoot)
	server.Handle("/api", HandleHome)
	err := server.Listen()
	isAnError(err)

}
