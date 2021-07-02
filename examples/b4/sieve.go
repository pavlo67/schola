package b4

import (
	"fmt"
	"math"
)

// version 1: empty boilerplate -----------------------------------------

func SieveV1(n int) []int {
	return []int{}
}

// version 2: structured boilerplate with stubs -------------------------

type sieve []bool

func SieveV2(n int) []int {
	//if n < 2 {
	//	return []int{}
	//}

	sieveInitial := InitSieveV2(n)
	sieveProcessed := ProcessNumbersFrom2ToSQRTNV2(sieveInitial)
	primesFromSieve := GetPrimesFromSieveV2(sieveProcessed)

	return primesFromSieve
}

func InitSieveV2(n int) sieve {
	return sieve{}
}

func ProcessNumbersFrom2ToSQRTNV2(sieveInitial sieve) sieve {
	return sieve{}
}

func GetPrimesFromSieveV2(sieveProcessed sieve) []int {
	return []int{}
}

// version 3: functions are partially implemented ----------------------

func SieveV3(n int) []int {
	//if n < 2 {
	//	return []int{}
	//}

	sieveInitial := InitSieveV3(n)
	sieveProcessed := ProcessNumbersFrom2ToSQRTNV3(sieveInitial)
	primesFromSieve := GetPrimesFromSieveV3(sieveProcessed)

	return primesFromSieve
}

func InitSieveV3(n int) sieve {
	if n < 2 {
		return sieve{}
	}

	// we use n-1 because we need to get 1-element array for the first number == 2 and so on
	sieveInitial := make(sieve, n-1)
	for i := range sieveInitial {
		sieveInitial[i] = true
	}

	return sieveInitial
}

func ProcessNumbersFrom2ToSQRTNV3(sieveInitial sieve) sieve {
	return sieveInitial
}

func GetPrimesFromSieveV3(sieveProcessed sieve) []int {
	fmt.Printf("final sieve to get primes: %#v\n\n", sieveProcessed)

	if len(sieveProcessed) == 1 {
		return []int{2}
	} else if len(sieveProcessed) == 2 {
		return []int{2, 3}
	}

	return []int{}
}

// version final: all functions are implemented -------------------------

func Sieve(n int) []int {
	//if n < 2 {
	//	return []int{}
	//}

	sieveInitial := InitSieve(n)
	sieveProcessed := ProcessNumbersFrom2ToSQRTN(sieveInitial)
	primesFromSieve := GetPrimesFromSieve(sieveProcessed)

	return primesFromSieve
}

func InitSieve(n int) sieve {
	if n < 2 {
		return sieve{}
	}

	// we use n-1 because we need to get 1-element array for the first number == 2 and so on
	sieveInitial := make(sieve, n-1)
	for i := range sieveInitial {
		sieveInitial[i] = true
	}

	return sieveInitial
}

func ProcessNumbersFrom2ToSQRTN(sieveToProcess sieve) sieve {
	// we could comment +1 here
	n := len(sieveToProcess) + 1
	if n < 2 {
		return sieveToProcess
	}
	nMaxToProcess := math.Sqrt(float64(n))

	for nToProcess := 2; float64(nToProcess) <= nMaxToProcess; nToProcess++ {
		// we could comment -2 here
		if !sieveToProcess[nToProcess-2] {
			continue
		}

		for nToSetComposite := nToProcess + nToProcess; nToSetComposite <= n; nToSetComposite += nToProcess {
			// we could comment -2 here
			sieveToProcess[nToSetComposite-2] = false
		}
	}

	return sieveToProcess
}

func GetPrimesFromSieve(sieveProcessed sieve) []int {
	primesFromSieve := []int{}

	for i := range sieveProcessed {
		if sieveProcessed[i] {
			// we use i+2 because the first i == 0 corresponds to the first prime number == 2 and so on
			primesFromSieve = append(primesFromSieve, i+2)
		}
	}

	return primesFromSieve
}
