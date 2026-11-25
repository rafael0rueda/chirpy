package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	var srv http.Server
	srv.Addr = ":8080"
	srv.Handler = mux

	mux.Handle("/", http.FileServer(http.Dir(".")))

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Server error: %v", err)
	}

}
