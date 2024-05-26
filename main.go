package main

import (
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(".")))

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Printf("Serving on port 8080\n")

	log.Fatal(srv.ListenAndServe())
}
