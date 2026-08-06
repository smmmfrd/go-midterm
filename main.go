package main

import (
	"fmt"
	"sync"
	"time"

	readfile "github.com/smmmfrd/go-midterm/1_read_file"
	fetch "github.com/smmmfrd/go-midterm/2_fetch"
	redis "github.com/smmmfrd/go-midterm/3_redis"
	jsonloader "github.com/smmmfrd/go-midterm/4_json_loader"
	ratelimiter "github.com/smmmfrd/go-midterm/5_rate_limiter"
)

func main() {
	var serverDuration time.Duration = 3 * time.Second
	var wg sync.WaitGroup

	printTitle("CREATING FILE OF RANDOM NUMBERS\nTHEN CALCULATING AVERAGE, MEAN, AND MODE")
	readfile.WriteFile()

	readfile.ReadFile()

	printTitle("FETCHING URLs CONCURRENTLY AND HANDLING THEIR ERRORS")
	fetch.FetchURLs()

	printTitle("CUSTOM REDIS SERVER HANDLING CREATE, READ, AND DELETE")
	wg.Add(1)
	go redis.Start(serverDuration, &wg)

	redis.Client()

	wg.Wait()

	printTitle("READING AND VALIDATING MULTIPLE JSON FILES, COMPILING ALL ERRORS")
	jsonloader.ReadJson()

	printTitle("DEMONSTRATING RATE LIMITED HTTP SERVER")
	wg.Add(1)
	go ratelimiter.Start(serverDuration, &wg)

	ratelimiter.BeginClient()
	wg.Wait()
}

func printTitle(title string) {
	fmt.Println("\n\n===============================================")
	fmt.Printf("%s\n", title)
	fmt.Println("===============================================")
}
