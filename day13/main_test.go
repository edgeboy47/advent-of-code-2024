package main

import "testing"

func TestMin(t *testing.T) {
	t.Run("Test Min", func(t *testing.T) {
		a := Button{94, 34}
		b := Button{22, 67}
		prize := Button{8400, 5400}
		cache := []Combination{}
		expected := Combination{80, 40}
		actual := min(a, b, prize, Combination{0, 0}, &cache)

		if expected != actual {
			t.Errorf("Expected %v, got %v", expected, actual)
		}
	})
	t.Run("Test Min no solution", func(t *testing.T) {
		a := Button{26, 66}
		b := Button{67, 21}
		prize := Button{12748, 12176}
		cache := []Combination{}
		expected := Combination{0, 0}
		actual := min(a, b, prize, Combination{0, 0}, &cache)

		if expected != actual {
			t.Errorf("Expected %v, got %v", expected, actual)
		}
	})
}
