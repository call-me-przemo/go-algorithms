package src

func QuickSort(numbers []int) {
	if len(numbers) == 1 {
		return
	}

	pivot := numbers[len(numbers)/2]
	i, j := 0, len(numbers)-1

	for i < j {
		if numbers[i] > pivot && numbers[j] <= pivot {
			numbers[i], numbers[j] = numbers[j], numbers[i]
			i++
			j--
		} else if numbers[i] <= pivot {
			i++
		} else {
			j--
		}

	}

	QuickSort(numbers[:i])
	QuickSort(numbers[i:])
}
