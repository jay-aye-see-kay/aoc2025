package aoc2025

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_all(t *testing.T) {
	tests := []struct {
		name     string
		part     int
		fn       func([]byte) (string, error)
		expected string
	}{
		{"day1_sample", 1, day1part1, "3"},
		{"day1", 1, day1part1, "964"},
		{"day1_sample", 2, day1part2, "6"},
		{"day1", 2, day1part2, "5872"},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s part %d", tt.name, tt.part), func(t *testing.T) {
			input, err := os.ReadFile("./inputs/" + tt.name + ".txt")
			require.NoError(t, err)
			result, err := tt.fn(input)
			require.NoError(t, err)
			assert.Equal(t, result, tt.expected)
		})
	}
}
