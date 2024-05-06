package task1

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTask1(t *testing.T) {
	cases := []struct {
		name     string
		in       int
		expected string
	}{
		{
			name:     "25",
			in:       25,
			expected: "25 компьютеров",
		},
		{
			name:     "41",
			in:       41,
			expected: "41 компьютер",
		},
		{
			name:     "1048",
			in:       1048,
			expected: "1048 компьютеров",
		},
		{
			name:     "5",
			in:       5,
			expected: "5 компьютеров",
		},
	}
	for _, tCase := range cases {
		t.Run(tCase.name, func(t *testing.T) {
			str := Editor(tCase.in)
			require.Equal(t, tCase.expected, str)
		})
	}
}
