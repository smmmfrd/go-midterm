package ratelimiter

import (
	"fmt"
	"io"
	"net/http"
)

func BeginClient() {
	fmt.Println("beginning spamming")
	resp, err := http.Get("http://localhost:9009/")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer resp.Body.Close()

	text, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Client received: %s\n", string(text))
}
