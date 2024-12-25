package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
)

func sumMulInstructions(input string) int {
	sum := 0

	regex, _ := regexp.Compile(`mul\((\d+),(\d+)\)`)
	matches := regex.FindAllStringSubmatch(input, -1)

	for _, match := range matches {
		arg1, _ := strconv.Atoi(match[1])
		arg2, _ := strconv.Atoi(match[2])

		sum += arg1 * arg2
	}
	return sum
}

func sumMulInstructionsWithCondition(input string, enabled bool) (int, bool) {
	sum := 0
  localEnabled := enabled

	regex, _ := regexp.Compile(`mul\((\d+),(\d+)\)|(do\(\))|(don't\(\))`)
	matches := regex.FindAllStringSubmatch(input, -1)

	for _, match := range matches {
    if slices.Contains(match, "do()") {
      localEnabled = true
      continue
    }

    if slices.Contains(match, "don't()") {
      localEnabled =  false
      continue
    }

    if !localEnabled {
      continue
    }

		arg1, _ := strconv.Atoi(match[1])
		arg2, _ := strconv.Atoi(match[2])

		sum += arg1 * arg2
	}
	return sum, localEnabled
}

func main() {
	inputFile := "input.txt"
	file, err := os.Open(inputFile)

	if err != nil {
		println("err reading file", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
  sum := 0
  enabled := true

  // Part 1
	for scanner.Scan() {
		line := scanner.Text()
    lineSum, newEnabled := sumMulInstructionsWithCondition(line, enabled)
    sum += lineSum
    enabled = newEnabled
	}

  fmt.Printf("Total sum: %d", sum)
}
