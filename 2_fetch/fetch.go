package fetch

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"
)

func FetchURLs() {
	fmt.Println("\n----- FETCHING SOME URLS CONCURRENTLY -----")

	data, err := os.ReadFile("data/urls.txt")
	if err != nil {
		fmt.Printf("Error reading file: %s\n", err.Error())
		return
	}

	urls := strings.Split(string(data), "\n")

	var wg sync.WaitGroup
	ch := make(chan int)

	for i, url := range urls {
		fmt.Printf("%d: %s\n", i, url)

		wg.Add(1)

		wg.Go(func() {
			fetch(url, ch, &wg)
		})
	}

	fmt.Println("We've starting the goroutines!")

	go func() {
		wg.Wait()
		close(ch)
	}()

	for result := range ch {
		fmt.Printf("%d\n", result)
	}

	fmt.Println("All done!")
}

func fetch(targetURL string, data chan int, wg *sync.WaitGroup) {
	defer wg.Done()

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

	data <- res.StatusCode
}
