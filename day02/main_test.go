package main

import "testing"

func TestIsStrictlyIncreasing(t *testing.T) {
	input := []int{1, 3, 6, 7, 9}
	expected := true
	var actual bool

	actual = isStrictlyIncreasing(input)

	if expected != actual {
		t.Errorf("Input is not strictly increasing")
	}
}

func TestIsStrictlyDecreasing(t *testing.T) {
	input := []int{7, 6, 4, 2, 1}
	expected := true
	var actual bool

	actual = isStrictlyDecreasing(input)

	if expected != actual {
		t.Errorf("Input is not strictly decreasing")
	}
}

func TestProblemDampener(t *testing.T) {
	input := []int{1, 3, 2, 4, 5}
	expected := true
	actual := problemDampener(input)

	if expected != actual {
		t.Errorf("Problem Dampener is not working")
	}
}
