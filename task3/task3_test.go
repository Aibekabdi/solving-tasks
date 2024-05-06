package task3

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTask3(t *testing.T) {
	cases := []struct {
		name     string
		in       [2]int
		expected []int
	}{
		{
			name:     "1",
			in:       [2]int{11, 20},
			expected: []int{11, 13, 17, 19},
		},
		{
			name:     "empty",
			in:       [2]int{11, -1},
			expected: []int{},
		},
		{
			name:     "2",
			in:       [2]int{1, 100},
			expected: []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97},
		},
	}
	for _, tCase := range cases {
		t.Run(tCase.name, func(t *testing.T) {
			str := primeNumbers(tCase.in[0], tCase.in[1])
			require.Equal(t, tCase.expected, str)
		})
	}
}
