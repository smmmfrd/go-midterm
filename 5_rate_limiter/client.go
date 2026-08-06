package ratelimiter

import (
	"fmt"
	"io"
	"net/http"
)

func BeginClient() {
	fmt.Println("beginning spamming")

	for i := range 10 {
		err := makeReq()
		if err != nil {
			fmt.Printf("Client failed after %d requests", i+1)
			break
		}
	}
}

func makeReq() error {
	resp, err := http.Get("http://localhost:9009/")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	text, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Printf("Client received: %s\n", string(text))
	return nil
}
