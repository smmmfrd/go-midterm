package ratelimiter

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

func Start(duration time.Duration) {
	fmt.Println("hello from the server with a rate limiter!")

	maxRequests := 5
	port := "9009"

	// Was thinking, have this make a server, then close it after ten seconds
	// Another script can spam it
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", limiterMiddleware(index, maxRequests))

	server := &http.Server{
		Addr:    "localhost" + ":" + port,
		Handler: mux,
	}

	go func() {
		log.Printf("Starting server on port %v\n", port)
		log.Fatal(server.ListenAndServe())
	}()

	time.Sleep(duration)

	fmt.Println("Shutting down")
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}

func limiterMiddleware(next http.HandlerFunc, maxRequests int) http.HandlerFunc {
	history := make(map[string]int)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		portIndex := strings.Index(r.RemoteAddr, ":")
		address := r.RemoteAddr[:portIndex]

		history[address] += 1

		if history[address] > maxRequests {
			http.Error(w, "Too many requests", http.StatusTooManyRequests)
			return
		}

		fmt.Printf("Received request from: %s. They have made %d requests.\n", address, history[address])
		next.ServeHTTP(w, r)
	})
}
