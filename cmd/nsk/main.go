package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Print("Starting application...")

	port := os.Getenv("PORT")
	if len(port) == 0 {
		log.Fatal("Port is unset")
	}

	s := http.Server{
		Addr:    ":" + port,
		Handler: nil,
	}
	http.HandleFunc("/", handler)

	err := s.ListenAndServe()
	if err != nil {
		log.Fatalf("Serever is stoped with error: %v", err)
	}

	///
	//s.Shutdown

	log.Print("Application has been stoped")
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, http.StatusText(http.StatusOK))
}
