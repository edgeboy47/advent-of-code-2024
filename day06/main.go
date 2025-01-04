package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

type Board [][]string

type Position struct {
	x int
	y int
}

type Guard struct {
	position  Position
	direction string
}

// Predict the guard's path on the board
func predictGuardPath(board Board, guard Guard) int {
	visited := []string{}
	done := false
	boardRows := len(board)
	boardCols := len(board[0])

	for !done {
		visitedString := fmt.Sprintf("%d,%d", guard.position.x, guard.position.y)
		if !slices.Contains(visited, visitedString) {
			visited = append(visited, visitedString)
		}

		// If the guard reaches the edge of the board, exit
		if (guard.position.x == 0 && guard.direction == "up") ||
			(guard.position.x == boardRows-1 && guard.direction == "down") ||
			(guard.position.y == 0 && guard.direction == "left") ||
			(guard.position.y == boardCols-1 && guard.direction == "right") {
			done = true
			continue
		}

		nextPosition := guard.position
		switch guard.direction {
		case "up":
			nextPosition.x -= 1
			break

		case "down":
			nextPosition.x += 1
			break

		case "left":
			nextPosition.y -= 1
			break

		case "right":
			nextPosition.y += 1
			break
		}

		if board[nextPosition.x][nextPosition.y] == "#" {
			// If the next position is an obstacle
			nextDirection := guard.direction
			switch guard.direction {
			case "up":
				nextDirection = "right"
				break

			case "down":
				nextDirection = "left"
				break

			case "left":
				nextDirection = "up"
				break

			case "right":
				nextDirection = "down"
				break
			}

			guard.direction = nextDirection
		} else {
			guard.position = nextPosition
		}
	}

	return len(visited)
}

func main() {
	inputFile := "input.txt"
	file, err := os.Open(inputFile)

	if err != nil {
		println("err reading file", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	board := Board{}
	guard := Guard{}

	for scanner.Scan() {
		line := scanner.Text()
		board = append(board, strings.Split(line, ""))
	}

	for index, row := range board {
		if slices.Contains(row, "^") {
			guard = Guard{
				position: Position{
					x: index,
					y: slices.Index(row, "^"),
				},
				direction: "up",
			}
		}
	}

  fmt.Printf("Guard visited %d unique positions\n", predictGuardPath(board, guard))
}