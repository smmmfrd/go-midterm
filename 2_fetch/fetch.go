package fetch

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"
)

type response struct {
	url  string
	code int
}

func FetchURLs() {
	data, err := os.ReadFile("data/urls.txt")
	if err != nil {
		fmt.Printf("Error reading file: %s\n", err.Error())
		return
	}

	urls := strings.Split(string(data), "\n")

	var wg sync.WaitGroup
	ch := make(chan response)

	for _, url := range urls {
		wg.Add(1)

		wg.Go(func() {
			defer wg.Done()
			fetch(url, ch)
		})
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for result := range ch {
		fmt.Printf("%s: %d\n", result.url, result.code)
	}
}

func fetch(targetURL string, data chan response) {
	res, err := http.Get(targetURL)
	if err != nil {
		fmt.Printf("Error occurred fetching URL: %s\n", err.Error())
		return
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 || err != nil {
		fmt.Printf("Failed fetch with code: %d and\nbody: %s\n", res.StatusCode, body)
		return
	}

	data <- response{
		url:  targetURL,
		code: res.StatusCode,
	}
}
