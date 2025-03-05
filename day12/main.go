package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

type Point struct {
	row, col int
}

type Region map[Point]int

type Garden struct {
	garden map[int][]Point
}

func (g *Garden) area(region int) int {
	points, exists := g.garden[region]

	if !exists {
		return 0
	}

	return len(points)
}

// Given a point, return the number of adjacent points with the same key
func (g *Garden) adjacentPoints(region int, point Point) int {
	adjacentPoints := 0
	keyPoints := g.garden[region]

	up := Point{point.row - 1, point.col}
	down := Point{point.row + 1, point.col}
	left := Point{point.row, point.col - 1}
	right := Point{point.row, point.col + 1}

	points := []Point{up, down, left, right}
	for _, p := range points {
		if slices.Contains(keyPoints, p) {
			adjacentPoints++
		}
	}

	return adjacentPoints
}

func (g *Garden) perimeter(region int) int {
	points, exists := g.garden[region]

	if !exists {
		return 0
	}

	total := 0

	for _, point := range points {
		total += 4 - g.adjacentPoints(region, point)
	}

	return total
}

// Find connected components using dfs
func dfs(board [][]string, visited map[Point]int, point Point, key string, region int) {
	_, seen := visited[point]
	if seen {
		return
	}

	if point.row < 0 || point.row >= len(board) || point.col < 0 || point.col >= len(board[0]) {
		return
	}

	// If the current point on the board has the passed in key, then it is in the same region
	if board[point.row][point.col] == key {
		visited[point] = region

		adjacent := []Point{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}

		// Check 4 adjacent points
		for _, p := range adjacent {
			dfs(board, visited, Point{point.row + p.row, point.col + p.col}, key, region)
		}
	}
}

func findRegions(board [][]string) Region {
	// Maps each point to its region number
	regions := make(map[Point]int)
	region := 0

	for row, line := range board {
		for col, key := range line {
			point := Point{row, col}
			_, seen := regions[point]
			if !seen {
				dfs(board, regions, point, key, region)
				region++

			}
		}
	}

	return regions
}

func main() {
	inputFile := "input.txt"
	file, err := os.Open(inputFile)

	if err != nil {
		println("err reading file", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	board := [][]string{}

	for scanner.Scan() {
		line := scanner.Text()
		board = append(board, strings.Split(line, ""))
	}

	regions := findRegions(board)

	g := Garden{
		garden: make(map[int][]Point),
	}

	for point, region := range regions {
		g.garden[region] = append(g.garden[region], point)
	}

	total := 0
	for key := range g.garden {
		area := g.area(key)
		perimeter := g.perimeter(key)
		total += area * perimeter
	}
	fmt.Printf("Total Score: %d\n", total)
}
