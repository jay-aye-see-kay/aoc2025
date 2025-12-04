package aoc2025

import (
	"strconv"
	"strings"
)

func day4part1(input []byte) (string, error) {
	grid := NewGrid(string(input))
	count := 0
	for cell, hasRoll := range grid.cells {
		if hasRoll && grid.countNeighbors(cell) <= 4 {
			count += 1
		}
	}
	return strconv.Itoa(count), nil
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
