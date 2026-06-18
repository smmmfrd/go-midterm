package readfile

import (
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
)

func WriteFile() {
	f, err := os.OpenFile("data/random.txt", os.O_WRONLY|os.O_TRUNC, 0644)
	check(err)

	defer f.Close()

	var data string
	for range 100 {
		data += strconv.Itoa(rand.IntN(10)) + "\n"
	}

	t, err := f.WriteString(data)
	check(err)

	fmt.Printf("Wrote the random string to file in %d bytes.\n", t)
}
