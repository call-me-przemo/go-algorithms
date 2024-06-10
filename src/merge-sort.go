package src

func MergeSort(numbers []int) []int {
	if len(numbers) == 1 {
		return numbers
	}

	idx := len(numbers) / 2
	a, b := MergeSort(numbers[:idx]), MergeSort(numbers[idx:])
	sorted := make([]int, len(numbers))

	for i, j, k := 0, 0, 0; i < len(numbers); i++ {
		if j < len(a) && k < len(b) {
			if a[j] < b[k] {
				sorted[i] = a[j]
				j++
			} else {
				sorted[i] = b[k]
				k++
			}
		} else if j < len(a) {
			sorted[i] = a[j]
			j++
		} else {
			sorted[i] = b[k]
			k++
		}
	}

	return sorted
}
