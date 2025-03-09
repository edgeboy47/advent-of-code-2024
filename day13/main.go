package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type Button struct {
	x, y int
}

type Combination struct {
	a, b int
}

func min(a, b, prize Button, combine Combination, cache *[]Combination) Combination {
	if slices.Contains(*cache, combine) {
		return Combination{0, 0}
	}

	totalA := combine.a
	totalB := combine.b
	curr := Button{a.x*totalA + b.x*totalB, a.y*totalA + b.y*totalB}

	if curr.x == prize.x && curr.y == prize.y {
		return Combination{totalA, totalB}
	}

	*cache = append(*cache, combine)
	if curr.x > prize.x || curr.y > prize.y {
		return Combination{0, 0}
	}

	aBtn := min(a, b, prize, Combination{totalA + 1, totalB}, cache)
	if aBtn.a != 0 || aBtn.b != 0 {
		return aBtn
	}

	bBtn := min(a, b, prize, Combination{totalA, totalB + 1}, cache)

	if bBtn.a != 0 || bBtn.b != 0 {
		return bBtn
	}

	return Combination{0, 0}
}

func main() {
	inputFile := "input.txt"
	file, err := os.Open(inputFile)

	if err != nil {
		println("err reading file", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	inputs := [][]Button{}
	currInput := []Button{}
	buttonRegex := regexp.MustCompile(`Button [A-Z]: X\+(\d+), Y\+(\d+)`)
	prizeRegex := regexp.MustCompile(`Prize: X=(\d+), Y=(\d+)`)

	curr := "a"
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			currInput = []Button{}
			continue
		}

		if curr == "a" || curr == "b" {
			matches := buttonRegex.FindStringSubmatch(line)
			x, _ := strconv.Atoi(matches[1])
			y, _ := strconv.Atoi(matches[2])

			btn := Button{x, y}
			currInput = append(currInput, btn)

			if curr == "a" {
				curr = "b"
			} else {
				curr = "prize"
			}
		} else if curr == "prize" {
			matches := prizeRegex.FindStringSubmatch(line)
			x, _ := strconv.Atoi(matches[1])
			y, _ := strconv.Atoi(matches[2])

			btn := Button{x, y}
			currInput = append(currInput, btn)

			inputs = append(inputs, currInput)
			curr = "a"
		}
	}

	total := 0

	for _, input := range inputs {
		cache := []Combination{}
		fmt.Printf("Checking %v\n", input)
		out := min(input[0], input[1], input[2], Combination{0, 0}, &cache)

		if !reflect.DeepEqual(out, Combination{0, 0}) {
			total += out.a*3 + out.b
		}
	}

	fmt.Println("total: ", total)
}
