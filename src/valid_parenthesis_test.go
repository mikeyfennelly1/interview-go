package src

import (
	"github.com/stretchr/testify/require"
	"testing"
)

type ValidParenthesesTestCases struct {
	s        string
	Expected bool
}

func TestValidParenthesis(t *testing.T) {
	cases := []ValidParenthesesTestCases{
		{
			s:        "()",
			Expected: true,
		},
		{
			s:        "()[]{}",
			Expected: true,
		},
		{
			s:        "(]",
			Expected: false,
		},
		{
			s:        "([])",
			Expected: true,
		},
	}

	for _, c := range cases {
		require.Equal(t, c.Expected, ValidParenthesis(c.s))
	}
}
