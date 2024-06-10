package src

func SelectionSort(numbers []int) {
	for i := 0; i < len(numbers); i++ {
		lowest := i
		for j := i; j < len(numbers); j++ {
			if numbers[j] < numbers[lowest] {
				lowest = j
			}
		}
		numbers[i], numbers[lowest] = numbers[lowest], numbers[i]
	}
}
