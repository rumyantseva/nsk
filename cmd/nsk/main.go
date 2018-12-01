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
	http.HandleFunc("/readyz", readyz)
	http.HandleFunc("/healthz", healthz)

	err := s.ListenAndServe()
	if err != nil {
		log.Fatalf("Serever is stoped with error: %v", err)
	}

	///
	//s.Shutdown

	log.Print("Application has been stoped")
}

func healthz(w http.ResponseWriter, r *http.Request) {
	log.Printf("HEALTH: Request received...")
	fmt.Fprint(w, http.StatusText(http.StatusOK))
}

func readyz(w http.ResponseWriter, r *http.Request) {
	log.Printf("READYZ: Request received...")
	fmt.Fprint(w, http.StatusText(http.StatusOK))
}
