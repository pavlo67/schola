package c1

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// doc-strings could be needed here
type TestCaseForSum struct {
	N           int
	ExpectedSum int
	ExpectedErr bool
}

var TestCasesForSum = []TestCaseForSum{
	{
		N:           -1,
		ExpectedSum: 0,
		ExpectedErr: true,
	},
	{
		N:           0,
		ExpectedSum: 0,
		ExpectedErr: false,
	},
	{
		N:           1,
		ExpectedSum: 1,
		ExpectedErr: false,
	},
	{
		N:           2,
		ExpectedSum: 3,
		ExpectedErr: false,
	},
	{
		N:           3,
		ExpectedSum: 6,
		ExpectedErr: false,
	},
	{
		N:           5,
		ExpectedSum: 15,
		ExpectedErr: false,
	},
}

func TestSumByLoop(t *testing.T) {
	for i, tc := range TestCasesForSum {
		t.Log(i)
		sum, err := SumByLoop(tc.N)
		require.Equal(t, tc.ExpectedSum, sum)
		if tc.ExpectedErr {
			require.Error(t, err)
		} else {
			require.NoError(t, err)
		}
	}
}

func TestSumByFormula(t *testing.T) {
	for i, tc := range TestCasesForSum {
		t.Log(i)
		sum, err := SumByFormula(tc.N)
		require.Equal(t, tc.ExpectedSum, sum)
		if tc.ExpectedErr {
			require.Error(t, err)
		} else {
			require.NoError(t, err)
		}
	}
}
