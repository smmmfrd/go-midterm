package jsonloader

import (
	"encoding/json"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type ServerJSON struct {
	Server struct {
		Host    string `json:"host"`
		Port    int    `json:"port"`
		Timeout int    `json:"timeout"`
	} `json:"server"`
	Database struct {
		Url            string `json:"url"`
		MaxConnections int    `json:"max_connections"`
	} `json:"database"`
	Logging struct {
		Level string `json:"level"`
		File  string `json:"file"`
	} `json:"logging"`
}

var files = []string{"data/json/good.json", "data/json/bad.json", "data/json/wrong.json"}

func ReadJson() {
	for _, v := range files {
		JsonLoader(v)
	}
}

func JsonLoader(filename string) {
	filedata, err := os.ReadFile(filename)
	check(err)

	fmt.Println(string(filedata))

	serverJSON := ServerJSON{}

	err = json.Unmarshal(filedata, &serverJSON)
	check(err)

	fmt.Println(serverJSON)
}
