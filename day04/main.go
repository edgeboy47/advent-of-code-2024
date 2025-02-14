package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// Searches the grid for the given word in 8 directions from the given coords
func searchPosition(grid [][]string, row, col int, word string) int {
	rows := len(grid)
	cols := len(grid[0])
	wordLength := len(word)
	count := 0

	// Skip if current letter is not the first letter of the word
	if grid[row][col] != string(word[0]) {
		return 0
	}

	// Directions to be searched
	x := []int{-1, -1, -1, 0, 0, 1, 1, 1}
	y := []int{-1, 0, 1, -1, 1, -1, 0, 1}

	// For each direction
	for dir := 0; dir < len(x); dir++ {
		currX := row + x[dir]
		currY := col + y[dir]
		index := 1

		// Check the word in current direction
		for index = 1; index < wordLength; index++ {
			// Check if indices are out of bounds
			if currX >= rows || currX < 0 || currY >= cols || currY < 0 {
				break
			}

			// If the current letter does not match the word, stop looking in this direction
			if grid[currX][currY] != string(word[index]) {
				break
			}

			currX += x[dir]
			currY += y[dir]
		}

		// If the index reached the length of the word, the word was found in this direction
		if index == wordLength {
			count++
		}
	}

	return count
}

func searchX(grid [][]string, row, col int) bool {
	rows := len(grid)
	cols := len(grid[0])

	if grid[row][col] != "A" || row == 0 || col == 0 || row == rows-1 || col == cols-1 {
		return false
	}

	topLeft := grid[row-1][col-1]
	topRight := grid[row-1][col+1]
	bottomLeft := grid[row+1][col-1]
	bottomRight := grid[row+1][col+1]

	diag1 := strings.Join([]string{topLeft, "A", bottomRight}, "")
	diag2 := strings.Join([]string{bottomLeft, "A", topRight}, "")
	regex, _ := regexp.Compile(`MAS|SAM`)

	return regex.MatchString(diag1) && regex.MatchString(diag2)
}

func searchGrid(grid [][]string, word string) int {
	rows := len(grid)
	cols := len(grid[0])
	count := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			found := searchX(grid, i, j)
			if found {
				count++
			}
		}
	}
	return count
}

func main() {
	inputFile := "input.txt"
	file, err := os.Open(inputFile)

	if err != nil {
		println("err reading file", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	input := [][]string{}
	word := "XMAS"
	row := 0

	// Part 1
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, strings.Split(line, ""))
		row++
	}

	found := searchGrid(input, word)
	fmt.Printf("found word %d times in input", found)
}
