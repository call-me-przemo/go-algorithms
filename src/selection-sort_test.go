package src

import (
	"slices"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	sl1 := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	sl2 := []int{5, 8, 9, 2, 0, 1}
	sl3 := []int{1, 2, 3, 4}

	SelectionSort(sl1)
	SelectionSort(sl2)
	SelectionSort(sl3)

	if !slices.Equal(sl1, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}) {
		t.Fatalf("Testing error")
	}

	if !slices.Equal(sl2, []int{0, 1, 2, 5, 8, 9}) {
		t.Fatalf("Testing error")
	}

	if !slices.Equal(sl3, []int{1, 2, 3, 4}) {
		t.Fatalf("Testing error")
	}
}
