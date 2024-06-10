package src

func BubbleSort(numbers []int) {
	loop := true

	j := len(numbers)
	for loop {
		loop = false

		for i := 1; i < j; i++ {
			if numbers[i] < numbers[i-1] {
				numbers[i], numbers[i-1] = numbers[i-1], numbers[i]
				loop = true
			}
		}
		j--
	}
}
