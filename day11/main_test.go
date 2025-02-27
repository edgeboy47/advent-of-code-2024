package main

import (
	"slices"
	"testing"
)

func TestBlink(t *testing.T) {
	t.Run("First Blink", func(t *testing.T) {
		input := []int{125, 17}
		expected := []int{253000, 1, 7}
		actual := blink(input)

		if !slices.Equal(expected, actual) {
			t.Errorf("Expected %v, got %v", expected, actual)
		}
	})

	t.Run("Length after 25 runs", func(t *testing.T) {
		input := []int{125, 17}
		expected := 55312
		actual := 0

		for range 25 {
			input = blink(input)
		}

		actual = len(input)

		if expected != actual {
			t.Errorf("Expected %d, got %d", expected, actual)
		}
	})
}
