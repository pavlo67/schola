package b4

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type TestCaseForPrimeNumbers struct {
	NotGreaterThan int
	PrimesExpected []int
}

var TestCasesForPrimeNumbers = []TestCaseForPrimeNumbers{
	{-5, []int{}},
	{-1, []int{}},
	{0, []int{}},
	{1, []int{}},
	{2, []int{2}},
	{3, []int{2, 3}},
	{5, []int{2, 3, 5}},
	{8, []int{2, 3, 5, 7}},
	{25, []int{2, 3, 5, 7, 11, 13, 17, 19, 23}},
	{55, []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53}},
	{75, []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73}},
}

// should fail (RED stage on RED-GREEN testing)
func TestSieveV1(t *testing.T) {
	for _, tc := range TestCasesForPrimeNumbers {
		t.Logf("testing n = %d", tc.NotGreaterThan)
		primesFound := SieveV1(tc.NotGreaterThan)
		require.Equal(t, tc.PrimesExpected, primesFound)
	}
}

// should fail (RED stage on RED-GREEN testing yet)
func TestSieveV2(t *testing.T) {
	for _, tc := range TestCasesForPrimeNumbers {
		t.Logf("testing n = %d", tc.NotGreaterThan)
		primesFound := SieveV2(tc.NotGreaterThan)
		require.Equal(t, tc.PrimesExpected, primesFound)
	}
}

// should fail (RED stage on RED-GREEN testing yet)
func TestSieveV3(t *testing.T) {
	for _, tc := range TestCasesForPrimeNumbers {
		t.Logf("testing n = %d", tc.NotGreaterThan)
		primesFound := SieveV3(tc.NotGreaterThan)
		require.Equal(t, tc.PrimesExpected, primesFound)
	}
}

// should pass (GREEN stage on RED-GREEN testing)
func TestSieve(t *testing.T) {
	for _, tc := range TestCasesForPrimeNumbers {
		t.Logf("testing n = %d", tc.NotGreaterThan)
		primesFound := Sieve(tc.NotGreaterThan)
		require.Equal(t, tc.PrimesExpected, primesFound)
	}
}
