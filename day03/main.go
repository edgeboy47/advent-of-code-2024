package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

func main() {
	inputFile := "input.txt"
	file, err := os.Open(inputFile)

	if err != nil {
		println("err reading file", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
  sum := 0

  // Part 1
	for scanner.Scan() {
		line := scanner.Text()
    sum += sumMulInstructions(line)
	}

  fmt.Printf("Total sum: %d", sum)
}
