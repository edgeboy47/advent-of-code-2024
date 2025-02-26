package main

import "testing"

func TestTrailHeadScore(t *testing.T) {
	topo := [][]int{
		{8, 9, 0, 1, 0, 1, 2, 3},
		{7, 8, 1, 2, 1, 8, 7, 4},
		{8, 7, 4, 3, 0, 9, 6, 5},
		{9, 6, 5, 4, 9, 8, 7, 4},
		{4, 5, 6, 7, 8, 9, 0, 3},
		{3, 2, 0, 1, 9, 0, 1, 2},
		{0, 1, 3, 2, 9, 8, 0, 1},
		{1, 0, 4, 5, 6, 7, 3, 2},
	}

	t.Run("Find score", func(t *testing.T) {
		expected := 5
		actual := trailHeadScore(topo, &[]string{}, 0, 2)

		if expected != actual {
			t.Errorf("Expected %d, got %d", expected, actual)
		}
	})

	t.Run("Boundary check", func(t *testing.T) {
		expected := 5
		actual := trailHeadScore(topo, &[]string{}, 7, 1)

		if expected != actual {
			t.Errorf("Expected %d, got %d", expected, actual)
		}
	})
}

func TestTrailHeadRating(t *testing.T) {
	topo := [][]int{
		{8, 9, 0, 1, 0, 1, 2, 3},
		{7, 8, 1, 2, 1, 8, 7, 4},
		{8, 7, 4, 3, 0, 9, 6, 5},
		{9, 6, 5, 4, 9, 8, 7, 4},
		{4, 5, 6, 7, 8, 9, 0, 3},
		{3, 2, 0, 1, 9, 0, 1, 2},
		{0, 1, 3, 2, 9, 8, 0, 1},
		{1, 0, 4, 5, 6, 7, 3, 2},
	}

	t.Run("Find score", func(t *testing.T) {
		expected := 20
		actual := trailHeadRating(topo, 0, 2)

		if expected != actual {
			t.Errorf("Expected %d, got %d", expected, actual)
		}
	})

	t.Run("Find score", func(t *testing.T) {
		expected := 24
		actual := trailHeadRating(topo, 0, 4)

		if expected != actual {
			t.Errorf("Expected %d, got %d", expected, actual)
		}
	})
}
