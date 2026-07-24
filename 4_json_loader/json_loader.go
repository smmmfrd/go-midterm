package jsonloader

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
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

	serverJSON := ServerJSON{}

	err = json.Unmarshal(filedata, &serverJSON)
	check(err)

	err = Verify(reflect.ValueOf(serverJSON))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

// Checks if each field has a value
func Verify(v reflect.Value) error {
	for i := range v.NumField() {
		field := v.Field(i)
		name := v.Type().Field(i).Name

		switch field.Kind() {
		case reflect.String:
			if len(field.String()) == 0 {
				return fmt.Errorf("Empty string found at %s", name)
			}
		}

		if v.Field(i).Kind() == reflect.Struct {
			return Verify(v.Field(i))
		}
	}

	return nil
}
