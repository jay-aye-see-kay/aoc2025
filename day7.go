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

// the idea here is any number >0 is a beam, every time a beam splits we increase
// (double?) it's numbers, so we keep track of "timelines" without having to
// duplicate the whole grid each split
func day7part2(input []byte) (int, error) {
	var (
		start    = -2
		splitter = -1

		grid = make([][]int, 0, bytes.Count(input, []byte{'\n'}))
	)

	// build a grid where negative numbers are special chars, postive are path counts
	for row := range bytes.Lines(input) {
		row = bytes.TrimSpace(row)
		newRow := make([]int, len(row))
		for i, col := range row {
			switch col {
			case byte('S'):
				newRow[i] = start
			case byte('^'):
				newRow[i] = splitter
			}
		}
		grid = append(grid, newRow)
	}

	// loop through grid (not last line) computing the next line's value
	for x, row := range grid[:len(grid)-1] {
		for y, col := range row {
			switch {
			case col == 0:
			// nothing
			case col > 0:
				if grid[x+1][y] != splitter {
					grid[x+1][y] = grid[x][y]
				}
			case col == splitter:
				inputCount := grid[x-1][y]
				if inputCount > 0 {
					grid[x][y-1] += inputCount
					grid[x][y+1] += inputCount
					grid[x+1][y-1] += inputCount
					grid[x+1][y+1] += inputCount
				}
			case col == start:
				grid[x+1][y] = 1
			}
		}
	}

	// sum path counts in last row
	count := 0
	for _, c := range grid[len(grid)-1] {
		count += c
	}
	return count, nil
}
