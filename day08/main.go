package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

type Frequencies map[string][]Coordinate
type Grid [][]string
type Coordinate struct {
	x, y int
}

func findAntinodes(a, b Coordinate, maxRows, maxCols int) []Coordinate {
	// Takes a pair of coordinates and returns the 2 antinodes
	antinodes := []Coordinate{}
	antinodes = append(antinodes, a, b)

	xDistance := a.x - b.x
	yDistance := a.y - b.y

	nodeA := Coordinate{a.x, a.y}
	nodeB := Coordinate{b.x, b.y}

	for nodeA.x < maxRows && nodeA.x >= 0 && nodeA.y < maxCols && nodeA.y >= 0 {
		nodeA = Coordinate{nodeA.x + xDistance, nodeA.y + yDistance}
		antinodes = append(antinodes, nodeA)
	}

	for nodeB.x < maxRows && nodeB.x >= 0 && nodeB.y < maxCols && nodeB.y >= 0 {
		nodeB = Coordinate{nodeB.x - xDistance, nodeB.y - yDistance}
		antinodes = append(antinodes, nodeB)
	}

	return antinodes
}

func countAntinodes(frequencies Frequencies, maxRows, maxCols int) int {
	// For each frequency in the map, go through the list of coordinates
	// For each pair of coordinates, find the 2 antinodes and add the coordinates to a list
	// Then go through the list of antinode coordinates and find the total unique number
	antinodes := []Coordinate{}

	for _, list := range frequencies {
		for i := 0; i < len(list); i++ {
			curr := list[i]
			for j := i + 1; j < len(list); j++ {
				currAntinodes := findAntinodes(curr, list[j], maxRows, maxCols)
				for _, node := range currAntinodes {
					if !slices.Contains(antinodes, node) &&
						node.x < maxRows && node.x >= 0 &&
						node.y < maxCols && node.y >= 0 {
						antinodes = append(antinodes, node)
					}
				}
			}
		}

	}

	return len(antinodes)
}

func main() {
	inputFile := "input.txt"
	// inputFile := "input_test.txt"
	file, err := os.Open(inputFile)

	if err != nil {
		println("err reading file", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	grid := Grid{}
	frequencies := Frequencies{}

	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, strings.Split(line, ""))
	}

	for rowIndex, row := range grid {
		for colIndex, val := range row {
			if val == "." {
				continue
			}

			key := val
			coords, exists := frequencies[key]

			if exists {
				frequencies[key] = append(coords, Coordinate{rowIndex, colIndex})
			} else {
				frequencies[key] = []Coordinate{{rowIndex, colIndex}}
			}
		}
	}

	fmt.Printf("Total unique antinodes: %d\n", countAntinodes(frequencies, len(grid), len(grid[0])))
}
