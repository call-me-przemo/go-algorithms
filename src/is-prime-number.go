package src

import "math"

func IsPrimeNumber(number int) bool {
	if number < 2 {
		return false
	}

	for i := 2; i <= int(math.Sqrt(float64(number))); i++ {
		if number%i == 0 {
			return false
		}
	}

	return true
}
