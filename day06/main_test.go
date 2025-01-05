package main

import (
	"strings"
	"testing"
)

func TestPredictGuardPath(t *testing.T) {
	t.Run("example input", func(t *testing.T) {
		input := [][]string{
			{"....#....."},
			{".........#"},
			{".........."},
			{"..#......."},
			{".......#.."},
			{".........."},
			{".#..^....."},
			{"........#."},
			{"#........."},
			{"......#..."},
		}

		for index, str := range input {
			input[index] = strings.Split(str[0], "")
		}

		guard := Guard{position: Position{x: 6, y: 4}, direction: "up"}

		expected := 41
		actual := predictGuardPath(input, guard)

		if expected != actual {
			t.Errorf("Error with predictGuardPath, expected %d got %d", expected, actual)
		}
	})

}

func TestPredictGuardLoop(t *testing.T) {
	t.Run("example input", func(t *testing.T) {
		input := [][]string{
			{"....#....."},
			{".........#"},
			{".........."},
			{"..#......."},
			{".......#.."},
			{".........."},
			{".#..^....."},
			{"........#."},
			{"#........."},
			{"......#..."},
		}

		for index, str := range input {
			input[index] = strings.Split(str[0], "")
		}

		input[6][3] = "#"
		guard := Guard{position: Position{x: 6, y: 4}, direction: "up"}

		expected := true
		actual := predictGuardLoop(input, guard)

		if expected != actual {
			t.Errorf("Error with predictGuardLoop")
		}
	})

	t.Run("Test number of loops", func(t *testing.T) {
		board := [][]string{
			{"....#....."},
			{".........#"},
			{".........."},
			{"..#......."},
			{".......#.."},
			{".........."},
			{".#..^....."},
			{"........#."},
			{"#........."},
			{"......#..."},
		}
		loops := 0
		expected := 6

		for index, str := range board {
			board[index] = strings.Split(str[0], "")
		}

		for i := 0; i < len(board); i++ {
			for j := 0; j < len(board[0]); j++ {
				currentPosition := board[i][j]
				if currentPosition == "#" || currentPosition == "^" {
					continue
				}

				newBoard := make([][]string, len(board))
				for i := range newBoard {
					newBoard[i] = make([]string, len(board[0]))
					copy(newBoard[i], board[i])
				}
				newBoard[i][j] = "#"
				guard := Guard{position: Position{x: 6, y: 4}, direction: "up"}
				if predictGuardLoop(newBoard, guard) {
					loops++
				} 
			}
		}

		if expected != loops {
			t.Errorf("Incorrect number of loops found, expected %d, got %d", expected, loops)
		}
	})
}
