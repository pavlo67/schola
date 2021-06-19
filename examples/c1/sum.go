package c1

import "fmt"

func SumByLoop(n int) (int, error) {
	if n < 0 {
		return 0, fmt.Errorf("wrong N (%d), should be non-negative", n)
	} else if n == 0 {
		return 0, nil
	}

	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}

	return sum, nil
}

func SumByFormula(n int) (int, error) {
	if n < 0 {
		return 0, fmt.Errorf("wrong N (%d), should be non-negative", n)
	}

	return n * (n + 1) / 2, nil
}
