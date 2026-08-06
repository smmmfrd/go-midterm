package ratelimiter

import (
	"fmt"
	"log"
	"net/http"
)

func Start() {
	fmt.Println("hello from the server with a rate limiter!")
	port := "9009"

	// Was thinking, have this make a server, then close it after ten seconds
	// Another script can spam it
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", index)

	server := &http.Server{
		Addr:    "localhost" + ":" + port,
		Handler: mux,
	}

	log.Printf("Starting server on port %v\n", port)
	log.Fatal(server.ListenAndServe())
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello\n"))
}
