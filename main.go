package main

import (
	"fmt"
	"math/rand"
	"time"

	"internal/swapSort"
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

func main() {
	sortedList := swapSort.SwapSort([]int{4, 2, 1, 5, 2, 8, 3, -13})
	fmt.Println(sortedList)

	sortedList2 := swapSort.SwapSort(generateRandomNumbers(1000))
	fmt.Println(sortedList2)
}
