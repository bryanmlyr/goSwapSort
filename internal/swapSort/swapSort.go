package swapSort

import (
	"fmt"
	"math/rand"
	"time"
)

func generateRandomNumbers(length int) []int {
	var slice []int

	for x := 0; x < length; x++ {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		slice = append(slice, r1.Intn(1000000))
	}

	return slice
}

func swap(value1 *int, value2 *int) {
	var tmp int = *value1

	*value1 = *value2
	*value2 = tmp
}

func swapSort(values []int) []int {
	for x := 0; x < len(values)-1; x++ {
		for y := x + 1; y < len(values); y++ {
			if values[x] > values[y] {
				swap(&values[x], &values[y])
			}
		}
	}

	return values
}

func SortList(length uint, printable bool) {
	if printable {
		res := swapSort(generateRandomNumbers(int(length)))
		fmt.Println(res)
	} else {
		swapSort(generateRandomNumbers(int(length)))

	}
}
