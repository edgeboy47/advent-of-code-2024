package main

import "testing"

func TestSearchPosition(t *testing.T) {
	t.Run("Single row with word", func(t *testing.T) {
		input := [][]string{{"X", "M", "A", "S"}}
		expected := 1
		actual := searchPosition(input, 0, 0, "XMAS")

		if expected != actual {
			t.Errorf("searchPosition produced incorrect output")
		}
	})

	t.Run("Single column with word", func(t *testing.T) {
		input := [][]string{{"X"}, {"M"}, {"A"}, {"S"}}
		expected := 1
		actual := searchPosition(input, 0, 0, "XMAS")

		if expected != actual {
			t.Errorf("searchPosition produced incorrect output")
		}
	})

	t.Run("Single row without word", func(t *testing.T) {
		input := [][]string{{"X", "M"}}
		expected := 0
		actual := searchPosition(input, 0, 0, "XMAS")

		if expected != actual {
			t.Errorf("searchPosition produced incorrect output")
		}
	})

	t.Run("Single column without word", func(t *testing.T) {
		input := [][]string{{"X"}, {"M"}}
		expected := 0
		actual := searchPosition(input, 0, 0, "XMAS")

		if expected != actual {
			t.Errorf("searchPosition produced incorrect output")
		}
	})
}

func TestSearchGrid(t *testing.T) {
	t.Run("grid with word", func(t *testing.T) {
		input := [][]string{
			{".", ".", "X", ".", ".", "."},
			{".", "S", "A", "M", "X", "."},
			{".", "A", ".", ".", "A", "."},
			{"X", "M", "A", "S", ".", "S"},
			{".", "X", ".", ".", ".", "."},
		}
		expected := 4
		actual := searchGrid(input, "XMAS")

		if expected != actual {
			t.Errorf("searchGrid produced incorrect output: got %d, expected %d", actual, expected)
		}
	})

	t.Run("example input", func(t *testing.T) {
		input := [][]string{
			{"M", "M", "M", "S", "X", "X", "M", "A", "S", "M"},
			{"M", "S", "A", "M", "X", "M", "S", "M", "S", "A"},
			{"A", "M", "X", "S", "X", "M", "A", "A", "M", "M"},
			{"M", "S", "A", "M", "A", "S", "M", "S", "M", "X"},
			{"X", "M", "A", "S", "A", "M", "X", "A", "M", "M"},
			{"X", "X", "A", "M", "M", "X", "X", "A", "M", "A"},
			{"S", "M", "S", "M", "S", "A", "S", "X", "S", "S"},
			{"S", "A", "X", "A", "M", "A", "S", "A", "A", "A"},
			{"M", "A", "M", "M", "M", "X", "M", "M", "M", "M"},
			{"M", "X", "M", "X", "A", "X", "M", "A", "S", "X"},
		}
		expected := 18
		actual := searchGrid(input, "XMAS")

		if expected != actual {
			t.Errorf("searchGrid produced incorrect output: got %d, expected %d", actual, expected)
		}
	})
}
