package aoc2025

import (
	"bytes"
)

func day7part1(input []byte) (int, error) {
	var (
		start    = byte('S')
		splitter = byte('^')
		beam     = byte('|')

		splitCount = 0
		grid       = [][]byte{}
	)

	for row := range bytes.Lines(input) {
		grid = append(grid, bytes.TrimSpace(row))
	}

	for x, row := range grid {
		for y, col := range row {
			switch col {
			case start:
				grid[x+1][y] = beam
			case splitter:
				if grid[x-1][y] == beam {
					grid[x][y-1] = beam
					grid[x][y+1] = beam
					grid[x+1][y-1] = beam
					splitCount += 1
				}
			case beam:
				if x < len(row) && y < len(grid) && grid[x+1][y] != splitter {
					grid[x+1][y] = beam
				}
			}
		}
	}
	return splitCount, nil
}
