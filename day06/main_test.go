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
