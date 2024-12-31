package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Rules map[int][]int

func isXBeforeY(rules Rules, x, y int) bool {
	currentRules, exists := rules[x]

	if !exists {
		return true
	}

	if slices.Contains(currentRules, y) {
		return false
	}

	return true
}

func isUpdateValid(rules Rules, line string) bool {
	input := strings.Split(line, ",")
	vals := []int{}

	for _, val := range input {
		num, _ := strconv.Atoi(val)
		vals = append(vals, num)
	}
	itemsLength := len(vals)

	for i := 0; i < itemsLength; i++ {
		curr := vals[i]
		for j := i + 1; j < itemsLength; j++ {
			if !isXBeforeY(rules, curr, vals[j]) {
				return false
			}
		}
	}

	return true
}

func main() {
	inputFile := "input.txt"
	file, err := os.Open(inputFile)
	gotRules := false
	rules := Rules{}

	if err != nil {
		println("err reading file", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	total := 0

	for scanner.Scan() {
		line := scanner.Text()
		if !gotRules && line == "" {
			gotRules = true
			continue
		}

		if gotRules {
			// Check updates
			if isUpdateValid(rules, line) {
				items := strings.Split(line, ",")
				length := len(items)
				var index int

				if length%2 == 0 {
					index = length / 2
				} else {
					index = (length - 1) / 2
				}

				val, _ := strconv.Atoi(items[index])

				total += val
			}
		} else {
			// Get rules
			vals := strings.Split(line, "|")
			x, _ := strconv.Atoi(vals[0])
			y, _ := strconv.Atoi(vals[1])

			if _, exists := rules[y]; exists {
				rules[y] = append(rules[y], x)
			} else {
				rules[y] = []int{x}
			}
		}
	}

	fmt.Printf("middle page count: %d", total)
}
