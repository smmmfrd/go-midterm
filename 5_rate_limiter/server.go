package ratelimiter

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func Start() {
	fmt.Println("hello from the server with a rate limiter!")
	var serverDuration time.Duration = 10

	port := "9009"

	// Was thinking, have this make a server, then close it after ten seconds
	// Another script can spam it
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", index)

	server := &http.Server{
		Addr:    "localhost" + ":" + port,
		Handler: mux,
	}

	go func() {
		log.Printf("Starting server on port %v\n", port)
		log.Fatal(server.ListenAndServe())
	}()

	time.Sleep(serverDuration * time.Second)

	fmt.Println("Shutting down")
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello\n"))
}
