package main

import "testing"

func TestIsXBeforeY(t *testing.T) {
	rules := Rules{
		61: []int{97, 47, 75},
		47: []int{97, 75},
		29: []int{75, 97, 53, 61, 47},
		75: []int{97},
		53: []int{47, 75, 61, 97},
		13: []int{97, 61, 29, 47, 75, 53},
	}

	t.Run("x does not go before y", func(t *testing.T) {
		x := 75
		y := 97
		expected := false
		actual := isXBeforeY(rules, x, y)

		if expected != actual {
			t.Errorf("Error with isXBeforeY")
		}
	})

	t.Run("x goes before y", func(t *testing.T) {
		x := 75
		y := 47
		expected := true
		actual := isXBeforeY(rules, x, y)

		if expected != actual {
			t.Errorf("Error with isXBeforeY")
		}
	})
}

func TestIsUpdateValid(t *testing.T) {
	rules := Rules{
		61: []int{97, 47, 75},
		47: []int{97, 75},
		29: []int{75, 97, 53, 61, 47},
		75: []int{97},
		53: []int{47, 75, 61, 97},
		13: []int{97, 61, 29, 47, 75, 53},
	}

	t.Run("line is valid", func(t *testing.T) {
		input := "75,47,61,53,29"
		expected := true
		actual := isUpdateValid(rules, input)

		if expected != actual {
			t.Errorf("Error with isUpdateValid")
		}
	})

	t.Run("line is valid", func(t *testing.T) {
		input := "97,61,53,29,13"
		expected := true
		actual := isUpdateValid(rules, input)

		if expected != actual {
			t.Errorf("Error with isUpdateValid")
		}
	})

	t.Run("line is invalid", func(t *testing.T) {
		input := "75,97,47,61,53"
		expected := false
		actual := isUpdateValid(rules, input)

		if expected != actual {
			t.Errorf("Error with isUpdateValid")
		}
	})

	t.Run("line is invalid", func(t *testing.T) {
		input := "61,13,29"
		expected := false
		actual := isUpdateValid(rules, input)

		if expected != actual {
			t.Errorf("Error with isUpdateValid")
		}
	})

	t.Run("line is invalid", func(t *testing.T) {
		input := "97,13,75,29,47"
		expected := false
		actual := isUpdateValid(rules, input)

		if expected != actual {
			t.Errorf("Error with isUpdateValid")
		}
	})
}
