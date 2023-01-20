package main

import (
	"log"
)

func main() {
	store, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}


	server := newAPIServer(":3000", store)
	server.Run()
}


