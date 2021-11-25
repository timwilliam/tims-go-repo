package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToBase(t *testing.T) {
	tests := []struct {
		number         int64
		targetBase     int
		name, expected string
	}{
		{
			name:       "120",
			number:     120,
			targetBase: 10,
			expected:   "120",
		},
		{
			name:       "438",
			number:     438,
			targetBase: 2,
			expected:   "ERROR",
		},
		{
			name:       "43981",
			number:     43981,
			targetBase: 15,
			expected:   "D071",
		},
		{
			name:       "03",
			number:     3,
			targetBase: 10,
			expected:   "3",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := ToBase(test.number, test.targetBase)
			assert.Equal(t, test.expected, result)
		})
	}
}
