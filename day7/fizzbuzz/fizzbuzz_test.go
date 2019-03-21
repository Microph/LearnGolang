package fizzbuzz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSay(t *testing.T) {
	testCases := []struct {
		n        int
		expected string
	}{
		{15, "FizzBuzz"},
		{3, "Fizz"},
		{5, "Buzz"},
		{2, "2"},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.expected, Say(tc.n))
	}
}
