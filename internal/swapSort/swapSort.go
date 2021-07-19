package swapSort

func swap(value1 *int, value2 *int) {
	var tmp int = *value1

	*value1 = *value2
	*value2 = tmp
}

func SwapSort(values []int) []int {
	for x := 0; x < len(values)-1; x++ {
		for y := x + 1; y < len(values); y++ {
			if values[x] > values[y] {
				swap(&values[x], &values[y])
			}
		}
	}

	return values
}
