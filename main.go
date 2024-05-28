package main

import (
	"log"
	"net/http"
)

type handler interface {
	serverHTTP(w http.ResponseWriter, r *http.Request)
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Add("content-type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(http.StatusText(http.StatusOK)))

	})
	mux.Handle("/app/*", http.FileServer(http.Dir(".")))

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Printf("Serving on port 8080\n")

	log.Fatal(srv.ListenAndServe())
}
