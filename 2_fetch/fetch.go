package fetch

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func FetchURLs() {
	fmt.Println("\n----- FETCHING SOME URLS CONCURRENTLY -----")

	data, err := os.ReadFile("data/urls.txt")
	if err != nil {
		fmt.Printf("Error reading file: %s\n", err.Error())
		return
	}

	urls := strings.Split(string(data), "\n")

	for i, url := range urls {
		fmt.Printf("%d: %s\n", i, url)
		fetch(url)
	}
}

func fetch(targetURL string) {
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

	fmt.Printf("%d\n", res.StatusCode)
}
