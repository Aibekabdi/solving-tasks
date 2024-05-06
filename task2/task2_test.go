package task2

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTask2(t *testing.T) {
	cases := []struct {
		name     string
		in       []int
		expected []int
	}{
		{
			name:     "1",
			in:       []int{42, 12, 18},
			expected: []int{2, 3, 6},
		},
		{
			name:     "2",
			in:       []int{6, 90, 12, 18, 30, 18},
			expected: []int{2, 3, 6},
		},
		{
			name:     "3",
			in:       []int{1, 2, 3, 4, 5},
			expected: []int{},
		},
	}
	for _, tCase := range cases {
		t.Run(tCase.name, func(t *testing.T) {
			str := commonDivisors(tCase.in)
			require.Equal(t, tCase.expected, str)
		})
	}
}
