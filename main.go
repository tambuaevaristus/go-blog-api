package main

import (
	"example/go-blog-api/internal/todo"
	"example/go-blog-api/internal/transport"
	"log"
)

func main() {

	svc := todo.NewService()
	server := transport.NewServer(svc)

	if err := server.Serve(); err != nil {
		log.Fatal(err)
	}
}
