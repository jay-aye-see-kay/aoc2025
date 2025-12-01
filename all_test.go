package aoc2025

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_all(t *testing.T) {
	tests := []struct {
		name     string
		fn       func([]byte) (string, error)
		expected string
	}{
		{"day1_sample", day1part1, "3"},
		{"day1", day1part1, "964"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input, err := os.ReadFile("./inputs/" + tt.name + ".txt")
			require.NoError(t, err)
			result, err := tt.fn(input)
			require.NoError(t, err)
			assert.Equal(t, result, tt.expected)
		})
	}
}
