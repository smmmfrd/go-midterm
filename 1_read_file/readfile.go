package readfile

import (
	"fmt"
	"os"
	"slices"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadFile() {
	fmt.Println("\n >>>>> READING DATA FROM FILE <<<<<")

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

	fmt.Printf("Average value: %v\n", findAvg(nums))
	fmt.Printf("Mean value: %v\n", findMean(nums))
	fmt.Printf("Value with highest occurence: %v\n", findHighOcc(nums))
	fmt.Println()
}

func findAvg(nums []int) float32 {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	return float32(sum) / float32(len(nums))
}

func findMean(nums []int) float32 {
	if len(nums)%2 == 0 {
		return (float32(nums[(len(nums)/2)-1]) + float32(nums[len(nums)/2])) / 2
	} else {
		return float32(nums[len(nums)/2])
	}
}

func findHighOcc(nums []int) int {
	count := make(map[int]int)

	for _, num := range nums {
		count[num] += 1
	}

	high := -1
	occ := 0
	for key, value := range count {
		if value > occ {
			high = key
			occ = value
		}
	}

	return high
}
