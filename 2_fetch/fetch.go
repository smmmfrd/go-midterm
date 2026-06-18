package fetch

import (
	"fmt"
	"io"
	"net/http"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func FetchURLs() {
	res, err := http.Get("http://www.google.com/robots.txt")
	check(err)

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		fmt.Printf("Failed fetch with code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	check(err)

	fmt.Printf("%s\n", body)
}
