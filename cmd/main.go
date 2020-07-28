package main

import (
	benchApi "bench-api"
	"log"
	"net/http"
)

func main() {
	server := benchApi.NewServer()

	log.Println("listening to requests on port 8080")
	err := http.ListenAndServe(":8080", server)
	if err != nil {
		log.Fatal(err)
	}
}
