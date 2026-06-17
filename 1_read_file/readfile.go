package readfile

import (
	"fmt"
	"os"
	"slices"
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

	slices.Sort(nums)

	count := make(map[int]int)
	sum := 0

	for _, num := range nums {
		count[num] += 1
		sum += num
	}

	average := float32(sum) / float32(len(nums))
	var mean float32

	if len(nums)%2 == 0 {
		mean = (float32(nums[(len(nums)/2)-1]) + float32(nums[len(nums)/2])) / 2
	} else {
		mean = float32(nums[len(nums)/2])
	}

	fmt.Printf("Average value: %v\n", average)
	fmt.Printf("Mean value: %v\n", mean)

	fmt.Println(count)
}
