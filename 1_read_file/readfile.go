package readfile

import (
	"fmt"
	"os"
)

func ReadFile() {
	// f, err := os.OpenFile(, os.O_WRONLY|os.O_TRUNC, 0644)
	// check(err)

	// defer f.Close()

	data, err := os.ReadFile("data/random.txt")
	check(err)

	var nums []int
	for _, b := range data {
		if b == 10 {
			continue
		}
		nums = append(nums, int(b)-48)
	}

	fmt.Println(nums)
}
