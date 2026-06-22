package fetch

import (
	"fmt"
	"io"
	"net/http"
)

func FetchURLs() {
	fmt.Println("\n----- FETCHING SOME URLS CONCURRENTLY -----")
	res, err := http.Get("https://this-does-not-exist-abc123.com")
	if err != nil {
		fmt.Printf("Error occurred: %s\n", err.Error())
		return
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 || err != nil {
		fmt.Printf("Failed fetch with code: %d and\nbody: %s\n", res.StatusCode, body)
		return
	}

	fmt.Printf("%s\n", body)
}
