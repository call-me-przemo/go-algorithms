package src

import (
	"testing"
)

func TestIsPrimeNumber(t *testing.T) {
	primes := [...]int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31}
	notPrimes := [...]int{1, 4, 6, 8, 9, 10, 12, 14, 15, 16, 18, 20}

	for _, v := range primes {
		if !IsPrimeNumber(v) {
			t.Fatalf("Testing error")
		}
	}

	for _, v := range notPrimes {
		if IsPrimeNumber(v) {
			t.Fatalf("Testing error")
		}
	}
}
