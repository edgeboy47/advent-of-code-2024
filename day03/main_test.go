package main

import "testing"

func TestSumMulInstructions(t *testing.T) {
	input := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	actual := sumMulInstructions(input)

	expected := 161

	if expected != actual {
		t.Errorf("Incorrect sum of mul instructions: %d", actual)
	}
}

func TestSumMulInstructionsWithCondition(t *testing.T) {
	input := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
	actual, enabled := sumMulInstructionsWithCondition(input, true)
	expected := 48

	if expected != actual || enabled != true {
		t.Errorf("Incorrect sum of mul instructions with conditions: %d", actual)
	}
}
