package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func blink(input []int) []int {
	output := []int{}

	for _, val := range input {
		if val == 0 {
			output = append(output, 1)
		} else if len(fmt.Sprintf("%d", val))%2 == 0 {
			str := fmt.Sprintf("%d", val)
			left := str[0 : len(str)/2]
			right := str[len(str)/2:]

			leftVal, _ := strconv.Atoi(left)
			rightVal, _ := strconv.Atoi(right)

			output = append(output, leftVal, rightVal)
		} else {
			output = append(output, val*2024)
		}
	}

	return output
}

func main() {
	inputFile := "input.txt"
	file, err := os.Open(inputFile)

	if err != nil {
		println("err reading file", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	input := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Split(line, " ")
		for _, val := range values {
			num, _ := strconv.Atoi(val)
			input = append(input, num)
		}
	}

	for range 25 {
		input = blink(input)
	}

	fmt.Printf("Length after 25 runs: %d", len(input))
}
