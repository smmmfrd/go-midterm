package main

import (
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
	readfile.WriteFile()

	readfile.ReadFile()

	fetch.FetchURLs()

	wg.Add(1)
	go redis.Start(serverDuration, &wg)

	redis.Client()

	wg.Wait()

	jsonloader.ReadJson()

	wg.Add(1)
	go ratelimiter.Start(serverDuration, &wg)

	ratelimiter.BeginClient()
	wg.Wait()
}
