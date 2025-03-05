package main

import "testing"

func TestGarden(t *testing.T) {
	garden := Garden{
		garden: map[int][]Point{
			0: {{0, 0}, {0, 1}, {0, 2}, {0, 3}},
			1: {{1, 0}, {1, 1}, {2, 0}, {2, 1}},
			2: {{1, 2}, {2, 2}, {2, 3}, {3, 3}},
			3: {{1, 3}},
			4: {{3, 0}, {3, 1}, {3, 2}},
		},
	}
	t.Run("Test area", func(t *testing.T) {
		expected := 4
		actual := garden.area(0)

		if expected != actual {
			t.Errorf("Expected %d, got %d\n", expected, actual)
		}

	})

	t.Run("Test adjacent points", func(t *testing.T) {
		expected := 2
		actual := garden.adjacentPoints(0, Point{0, 1})

		if expected != actual {
			t.Errorf("Expected %d, got %d\n", expected, actual)
		}
	})

	t.Run("Test perimeter", func(t *testing.T) {
		expected := 10
		actual := garden.perimeter(2)

		if expected != actual {
			t.Errorf("Expected %d, got %d\n", expected, actual)
		}
	})
}
