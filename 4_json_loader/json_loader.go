package jsonloader

import (
	"fmt"
	"os"
	"reflect"
	"strings"
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

var files = []string{ /*"data/json/good.json", "data/json/bad.json",*/ "data/json/wrong.json"}

func ReadJson() {
	for _, v := range files {
		ValidateJSON(v)
	}
}

func ValidateJSON(filename string) {
	filedata, err := os.ReadFile(filename)
	check(err)

	replacer := strings.NewReplacer(
		" ", "",
		"\n", "",
		"\t", "",
		"{", "",
		"}", "",
		"\"", "",
		"://", ".//",
		"localhost:", "localhost.",
	)
	data := replacer.Replace(string(filedata))

	s := replacer.Replace(structToString(reflect.TypeFor[ServerJSON]()))

	compareJsonArrays(jsonStringToArray(data), jsonStringToArray(s))
}

func structToString(t reflect.Type) string {
	var keys []string
	for i := range t.NumField() {
		key := ""

		field := t.Field(i)

		if field.Type.Kind() == reflect.Struct {
			key += fmt.Sprintf("\"%s\":", field.Tag.Get("json"))
			key += structToString(field.Type)
		} else {
			key += fmt.Sprintf("\"%s\":%s", field.Tag.Get("json"), field.Type.String())
		}
		if i != t.NumField()-1 {
			key += ","
		}

		keys = append(keys, key)
	}

	return fmt.Sprintf("{%s}", strings.Join(keys, ""))
}

func jsonStringToArray(j string) []string {
	var arr []string
	caretLocation, currentWordIndex := 0, 0
	for {
		caretLocation += currentWordIndex
		currentWordIndex = strings.IndexFunc(j[caretLocation:], func(r rune) bool { return r == ':' || r == ',' })

		if currentWordIndex < 0 {
			arr = append(arr, j[caretLocation:])
			break
		}

		arr = append(arr, j[caretLocation:caretLocation+currentWordIndex])
		currentWordIndex++
	}
	return arr
}

func compareJsonArrays(data, shape []string) {
	var smaller, larger []string
	if len(data) > len(shape) {
		smaller = shape
		larger = data
	} else {
		smaller = data
		larger = shape
	}

	for i := 0; i < len(larger)-1; i++ {
		if i >= len(smaller) {
			fmt.Println("Smaller missing key: " + larger[i])
			break
		}

		if larger[i] != smaller[i] {
			fmt.Printf("Incorrect key in file: \n\tExpected: %s\n\tFound: %s\n", larger[i], smaller[i])
			break
		}

		// Check if the next value is a data type, if it is we skip that index
		switch larger[i+1] {
		case "string":
			fallthrough
		case "int":
			i++
		}
	}
}
