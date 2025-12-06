package aoc2025

import (
	"strings"
)

func day4part1(input []byte) (int, error) {
	grid := NewGrid(string(input))
	count := len(grid.findRemovable())
	return count, nil
}

func day4part2(input []byte) (int, error) {
	grid := NewGrid(string(input))
	initialCount := grid.count()

	count := initialCount
	for {
		// remove what can be
		for _, removable := range grid.findRemovable() {
			grid.cells[removable] = false
		}
		// stop once not more can be removed
		newCount := grid.count()
		if newCount < count {
			count = newCount
		} else {
			return initialCount - newCount, nil
		}
	}
}

type Coord struct {
	x int
	y int
}

func NewGrid(input string) *Grid {
	g := &Grid{
		cells: make(map[Coord]bool),
	}
	for y, line := range strings.Split(strings.TrimSpace(input), "\n") {
		for x, content := range strings.Split(strings.TrimSpace(line), "") {
			g.cells[Coord{x, y}] = content == "@"
		}
	}
	return g
}

type Grid struct {
	cells map[Coord]bool
}

func (g *Grid) count() int {
	totalRolls := 0
	for _, hasRoll := range g.cells {
		if hasRoll {
			totalRolls += 1
		}
	}
	return totalRolls
}

func (g *Grid) findRemovable() []Coord {
	removable := []Coord{}
	for cell, hasRoll := range g.cells {
		if hasRoll && g.countNeighbors(cell) <= 4 {
			removable = append(removable, cell)
		}
	}
	return removable
}

func (g *Grid) neighbors(cell Coord) []Coord {
	neighbors := []Coord{}
	for x := cell.x - 1; x <= cell.x+1; x++ {
		for y := cell.y - 1; y <= cell.y+1; y++ {
			neighbors = append(neighbors, Coord{x, y})
		}
	}

	return neighbors
}

func (g *Grid) countNeighbors(cell Coord) int {
	count := 0
	for _, cell := range g.neighbors(cell) {
		if g.cells[cell] {
			count += 1
		}
	}
	return count
}
