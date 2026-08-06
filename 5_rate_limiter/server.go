package ratelimiter

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

func Start() {
	fmt.Println("hello from the server with a rate limiter!")
	var serverDuration time.Duration = 10

	port := "9009"

	// Was thinking, have this make a server, then close it after ten seconds
	// Another script can spam it
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", limiterMiddleware(index))

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

func limiterMiddleware(next http.HandlerFunc) http.HandlerFunc {
	history := make(map[string]int)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		portIndex := strings.Index(r.RemoteAddr, ":")
		address := r.RemoteAddr[:portIndex]

		history[address] += 1

		fmt.Printf("Received request from: %s. They have made %d requests.\n", address, history[address])
		next.ServeHTTP(w, r)
	})
}
