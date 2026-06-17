package main

import (
	readfile "github.com/smmmfrd/go-midterm/1_read_file"
	fetch "github.com/smmmfrd/go-midterm/2_fetch"
)

func main() {
	// readfile.WriteFile()

	readfile.ReadFile()

	fetch.FetchURLs()
}
