package src

func SelectionSort(numbers []int) {
	for i := 0; i < len(numbers); i++ {
		min := i
		for j := i; j < len(numbers); j++ {
			if numbers[j] < numbers[min] {
				min = j
			}
		}
		numbers[i], numbers[min] = numbers[min], numbers[i]
	}
}
