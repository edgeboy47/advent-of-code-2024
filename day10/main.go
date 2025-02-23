package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func trailHeadScore(maps [][]int, peaks *[]string, row, col int) int {
	maxRows := len(maps) - 1
	maxCols := len(maps[0]) - 1

	if col > maxCols || row > maxRows || col < 0 || row < 0 {
		return 0
	}

	curr := maps[row][col]
	if curr == 9 {
		peak := fmt.Sprintf("%d,%d", row, col)
		if slices.Contains(*peaks, peak) {
			return 0
		}

		*peaks = append(*peaks, peak)
		return 1
	}

	score := 0
	if row > 0 {
		up := maps[row-1][col]

		if up == curr+1 {
			score += trailHeadScore(maps, peaks, row-1, col)
		}
	}

	if row < maxRows {
		down := maps[row+1][col]

		if down == curr+1 {
			score += trailHeadScore(maps, peaks, row+1, col)
		}
	}

	if col > 0 {
		left := maps[row][col-1]

		if left == curr+1 {
			score += trailHeadScore(maps, peaks, row, col-1)
		}

	}

	if col < maxCols {
		right := maps[row][col+1]

		if right == curr+1 {
			score += trailHeadScore(maps, peaks, row, col+1)
		}
	}

	return score
}

func main() {
	inputFile := "input.txt"
	file, err := os.Open(inputFile)

	if err != nil {
		println("err reading file", err)
	}
	defer file.Close()

	topoMap := [][]int{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		newLine := []int{}
		for _, val := range strings.Split(line, "") {
			conv, _ := strconv.Atoi(val)
			newLine = append(newLine, conv)
		}
		topoMap = append(topoMap, newLine)
	}
	scores := 0

	for row := range topoMap {
		for col := range topoMap[row] {
			if topoMap[row][col] == 0 {
				scores += trailHeadScore(topoMap, &[]string{}, row, col)
			}
		}
	}

	fmt.Printf("scores: %v", scores)
}
