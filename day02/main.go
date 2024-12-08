package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isStrictlyIncreasing(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
  prev := nums[0]

  for i := 1; i < len(nums); i++ {
    if (nums[i] <= prev) {
      return false
    }

    if (nums[i] - prev > 3) {
      return false
    }

    prev = nums[i]
  }

	return true
}

func isStrictlyDecreasing(nums []int) bool {
	if len(nums) == 0 {
		return false
	}

  prev := nums[0]

  for i := 1; i < len(nums); i++ {
    if (nums[i] >= prev) {
      return false
    }

    if (prev - nums[i] > 3) {
      return false
    }

    prev = nums[i]
  }

	return true
}

func main() {
	inputFile := "input.txt"
	file, err := os.Open(inputFile)

	if err != nil {
		println("err reading file", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	numSafe := 0

	// Part 1
	for scanner.Scan() {
		line := scanner.Text()
		vals := []int{}

		for _, val := range strings.Split(line, " ") {
			num, err := strconv.Atoi(val)
			if err != nil {
				fmt.Printf("Error converting num %s", err)
			} else {
				vals = append(vals, num)
			}
		}

		if isStrictlyIncreasing(vals) || isStrictlyDecreasing(vals) {
			numSafe++
		}
	}

	fmt.Printf("Safe Reports: %d", numSafe)
}
