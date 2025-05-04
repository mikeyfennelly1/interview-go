package src

import (
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

type TwoSumTestCases struct {
	Input    []int
	Target   int
	Expected []int
}

func TestTwoSum(t *testing.T) {
	cases := []TwoSumTestCases{
		{
			Input:    []int{2, 7, 11, 15},
			Target:   9,
			Expected: []int{0, 1},
		},
		{
			Input:    []int{3, 2, 4},
			Target:   6,
			Expected: []int{1, 2},
		},
	}

	for _, c := range cases {
		result := TwoSum(c.Input, c.Target)
		require.Equal(t, c.Expected, result)
		if reflect.DeepEqual(c.Expected, result) {
			logrus.Infof("%v == %v", c.Expected, result)
		}
	}
}
