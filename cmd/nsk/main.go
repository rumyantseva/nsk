package main

import (
	"log"
	"net/http"
)

func main() {
	log.Print("Starting application...")

	port := "8080"

	s := http.Server{
		Addr:    ":" + port,
		Handler: nil,
	}
	err := s.ListenAndServe()
	if err != nil {
		log.Fatalf("Serever is stoped with error: %v", err)
	}

	log.Print("Application has been stoped")
}
