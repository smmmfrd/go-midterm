package main

import (
	"time"

	redis "github.com/smmmfrd/go-midterm/3_redis"
)

func main() {
	// readfile.WriteFile()

	// readfile.ReadFile()

	// fetch.FetchURLs()

	go redis.Start()

	time.Sleep(10 * time.Millisecond)

	redis.Run()
}
