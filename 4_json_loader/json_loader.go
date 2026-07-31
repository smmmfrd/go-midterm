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
	)
	data := replacer.Replace(string(filedata))

	s := replacer.Replace(structToString(reflect.TypeFor[ServerJSON]()))

	validateDataToStructByString(data, s)
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

func validateDataToStructByString(d string, s string) {
	fmt.Println(d)
	fmt.Println(s)

	caretLocation, currentWordIndex := 0, 0
	for {
		caretLocation += currentWordIndex
		currentWordIndex = strings.IndexFunc(s[caretLocation:], func(r rune) bool { return r == ':' || r == ',' })

		if currentWordIndex < 0 {
			break
		}

		fmt.Println(s[caretLocation : caretLocation+currentWordIndex])
		currentWordIndex++
	}
}
