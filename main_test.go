package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateCaptcha(t *testing.T) {
	testCases := []struct {
		format       int
		leftOperand  int
		operator     int
		rightOperand int
		expected     string
	}{
		{
			format:       0,
			leftOperand:  1,
			operator:     0,
			rightOperand: 1,
			expected:     "one + 1",
		},
		{
			format:       1,
			leftOperand:  1,
			operator:     0,
			rightOperand: 1,
			expected:     "1 + one",
		},
		{
			format:       0,
			leftOperand:  1,
			operator:     1,
			rightOperand: 1,
			expected:     "one - 1",
		},
		{
			format:       1,
			leftOperand:  1,
			operator:     1,
			rightOperand: 1,
			expected:     "1 - one",
		},
		{
			format:       0,
			leftOperand:  1,
			operator:     2,
			rightOperand: 1,
			expected:     "one * 1",
		},
		{
			format:       1,
			leftOperand:  1,
			operator:     2,
			rightOperand: 1,
			expected:     "1 * one",
		},
	}

	for _, v := range testCases {
		actual := generateCaptcha(v.format, v.leftOperand, v.operator, v.rightOperand)

		assert.Equal(t, v.expected, actual)
	}
}
