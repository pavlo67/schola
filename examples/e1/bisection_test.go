package e1

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// doc-strings could be needed here
type TestCaseForSearchWithBisection struct {
	List     []int
	Value    int
	Expected bool
}

var TestCasesForSearchWithBisection = []TestCaseForSearchWithBisection{
	{List: []int{-2, 0, 1, 2, 7, 8, 11, 13}, Value: -4, Expected: false},
	{List: []int{-2, 0, 1, 2, 7, 8, 11, 13}, Value: -3, Expected: false},
	{List: []int{-2, 0, 1, 2, 7, 8, 11, 13}, Value: -2, Expected: true},
	{List: []int{-2, 0, 1, 2, 7, 8, 11, 13}, Value: -1, Expected: false},
	{List: []int{-2, 0, 1, 2, 7, 8, 11, 13}, Value: 0, Expected: true},
	{List: []int{-2, 0, 1, 2, 7, 8, 11, 13}, Value: 1, Expected: true},
	{List: []int{-2, 0, 1, 2, 7, 8, 11, 13}, Value: 2, Expected: true},
	{List: []int{-2, 0, 1, 2, 7, 8, 11, 13}, Value: 3, Expected: false},
	{List: []int{-2, 0, 1, 2, 7, 8, 11, 13}, Value: 4, Expected: false},
	{List: []int{-2, 0, 1, 2, 7, 8, 11, 13}, Value: 5, Expected: false},
	{List: []int{-2, 0, 1, 2, 7, 8, 11, 13}, Value: 6, Expected: false},
	{List: []int{-2, 0, 1, 2, 7, 8, 11, 13}, Value: 7, Expected: true},
	{List: []int{-2, 0, 1, 2, 7, 8, 11, 13}, Value: 8, Expected: true},
	{List: []int{-2, 0, 1, 2, 7, 8, 11, 13}, Value: 9, Expected: false},
	{List: []int{-2, 0, 1, 2, 7, 8, 11, 13}, Value: 10, Expected: false},
	{List: []int{-2, 0, 1, 2, 7, 8, 11, 13}, Value: 11, Expected: true},
	{List: []int{-2, 0, 1, 2, 7, 8, 11, 13}, Value: 12, Expected: false},
	{List: []int{-2, 0, 1, 2, 7, 8, 11, 13}, Value: 13, Expected: true},
	{List: []int{-2, 0, 1, 2, 7, 8, 11, 13}, Value: 14, Expected: false},
	{List: []int{-2, 0, 1, 2, 7, 8, 11, 13}, Value: 15, Expected: false},
	{List: []int{}, Value: -4, Expected: false},
	{List: nil, Value: -4, Expected: false},
}

func TestSearchWithBisection(t *testing.T) {
	for i, tc := range TestCasesForSearchWithBisection {
		t.Log(i)
		found := SearchWithBisection(tc.List, tc.Value)
		require.Equalf(t, tc.Expected, found, "searching %d", tc.Value)
	}
}
