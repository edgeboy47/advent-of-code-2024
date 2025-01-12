package main

import (
	// "reflect"
	"slices"
	"testing"
)

func nodesAreEqual(a, b []Coordinate) bool {
	for _, val := range a {
		if !slices.Contains(b, val) {
			return false
		}
	}

	return true
}

func TestFindAntinodes(t *testing.T) {
	t.Run("Find antinodes", func(t *testing.T) {
		testValues := []struct {
			a         Coordinate
			b         Coordinate
			antinodes []Coordinate
		}{
			{Coordinate{1, 8}, Coordinate{2, 5}, []Coordinate{{0, 11}, {3, 2}}},
			{Coordinate{2, 5}, Coordinate{1, 8}, []Coordinate{{0, 11}, {3, 2}}},
			{Coordinate{3, 7}, Coordinate{4, 4}, []Coordinate{{2, 10}, {5, 1}}},
			{Coordinate{4, 4}, Coordinate{3, 7}, []Coordinate{{2, 10}, {5, 1}}},
			{Coordinate{2, 5}, Coordinate{3, 7}, []Coordinate{{1, 3}, {4, 9}}},
			{Coordinate{3, 7}, Coordinate{2, 5}, []Coordinate{{1, 3}, {4, 9}}},
		}

		for _, testVals := range testValues {
			expected := testVals.antinodes
			actual := findAntinodes(testVals.a, testVals.b)

			if !nodesAreEqual(expected, actual) {
				t.Errorf("Error finding antinodes, got %d and %d", expected, actual)
			}
		}
	})
}
