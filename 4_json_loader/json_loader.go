package jsonloader

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var files = []string{"data/json/good.json", "data/json/bad.json"}

func ReadJson() {
	for _, v := range files {
		JsonLoader(v)
	}
}

func JsonLoader(filename string) {
	data, err := os.ReadFile(filename)
	check(err)

	fmt.Println(string(data))
}
