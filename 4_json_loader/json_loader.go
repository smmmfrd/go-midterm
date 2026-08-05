package jsonloader

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
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

var files = []string{"data/json/good.json", "data/json/bad.json", "data/json/wrong.json", "data/json/mispelled.json"}

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

	issues := compareJsonArrays(jsonStringToArray(data), jsonStringToArray(s))

	if len(issues) > 0 {
		fmt.Println("Issues were found with converting: " + filename)
		for _, v := range issues {
			fmt.Printf("\t%s\n", v)
		}
	} else {
		fmt.Println(filename + " was converted successfully.")
	}
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

func compareJsonArrays(data, shape []string) []string {
	var smaller, larger []string
	if len(data) > len(shape) {
		smaller = shape
		larger = data
	} else {
		smaller = data
		larger = shape
	}

	var issues []string
	for i := 0; i < len(larger)-1; i++ {
		if i >= len(smaller) {
			issues = append(issues, "Missing key: "+larger[i])
			break
		}

		if larger[i] != smaller[i] {
			issues = append(issues, fmt.Sprintf("Incorrect key, Expected: %s\tFound: %s", larger[i], smaller[i]))
			break
		}

		// Program specific evals
		switch larger[i] {
		case "port":
			portValue, err := strconv.Atoi(smaller[i+1])
			if err != nil {
				issues = append(issues, "Issue with port value from file: "+err.Error())
			}

			if portValue < 1 || portValue > 65535 {
				issues = append(issues, "Bad port value")
			}
		case "timeout":
			fallthrough
		case "max_connections":
			connections, err := strconv.Atoi(smaller[i+1])
			if err != nil {
				issues = append(issues, "Issue with max_connection value from file: "+err.Error())
			}

			if connections <= 0 {
				issues = append(issues, "Bad max connections value")
			}
		case "level":
			if !strings.Contains("infodebugwarnerror", smaller[i+1]) {
				issues = append(issues, "Bad logging level")
			}
		case "host":
			fallthrough
		case "url":
			if smaller[i+1] == "" {
				issues = append(issues, "No url specified for "+larger[i])
			}
		}

		// Check if the next value is a data type, if it is we skip that index
		switch larger[i+1] {
		case "string":
			fallthrough
		case "int":
			i++
		}
	}

	return issues
}
