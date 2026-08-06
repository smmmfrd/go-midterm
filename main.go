package main

import (
	"time"

	ratelimiter "github.com/smmmfrd/go-midterm/5_rate_limiter"
)

func main() {
	var serverDuration time.Duration = 10 * time.Second
	// readfile.WriteFile()

	// readfile.ReadFile()

	// fetch.FetchURLs()

	// go redis.Start()

	// time.Sleep(serverDuration * time.Second)

	// redis.Client()

	// jsonloader.ReadJson()

	go ratelimiter.Start(serverDuration)

	time.Sleep(serverDuration)
}
