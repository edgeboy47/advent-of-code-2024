package main

import (
	"maps"
	"testing"
)

func TestBlink(t *testing.T) {
	t.Run("First Blink", func(t *testing.T) {
		input := map[int]int{
			125: 1,
			17:  1,
		}
		expected := map[int]int{
			253000: 1,
			1:      1,
			7:      1,
		}
		cache := make(map[int][]int)
		actual := blink(input, cache)

		if !maps.Equal(expected, actual) {
			t.Errorf("Expected %v, got %v", expected, actual)
		}
	})

	t.Run("Length after 25 runs", func(t *testing.T) {
		input := map[int]int{
			125: 1,
			17:  1,
		}
		expected := 55312
		cache := make(map[int][]int)
		actual := 0

		for range 25 {
			input = blink(input, cache)
		}

		for _, val := range input {
			actual += val
		}

		if expected != actual {
			t.Errorf("Expected %d, got %d", expected, actual)
		}
	})
}
