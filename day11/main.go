package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func addToMap(input map[int]int, key int) map[int]int {
	if val, exists := input[key]; exists {
		input[key] = val + 1
	} else {
		input[key] = 1
	}
	return input
}

// The cache stores the number of stones returned for a given starting number after n blinks
func blinkRecursive(input, n int, cache map[string]int) int {
	if n == 0 {
		return 1
	}

	cacheKey := fmt.Sprintf("%d, %d", input, n)

	_, existsInCache := cache[cacheKey]

	if existsInCache {
		return cache[cacheKey]
	}

	result := 0

	if input == 0 {
		result = blinkRecursive(1, n-1, cache)
	} else if len(fmt.Sprintf("%d", input))%2 == 0 {
		str := fmt.Sprintf("%d", input)
		left := str[0 : len(str)/2]
		right := str[len(str)/2:]

		leftVal, _ := strconv.Atoi(left)
		rightVal, _ := strconv.Atoi(right)

		result += blinkRecursive(leftVal, n-1, cache)
		result += blinkRecursive(rightVal, n-1, cache)
	} else {
		result = blinkRecursive(input*2024, n-1, cache)
	}

	cache[cacheKey] = result

	return result
}

// Store the list as a map, where the key is the number and the value is the count of that number in the list
func blink(input map[int]int, cache map[int][]int) map[int]int {
	output := make(map[int]int)
	// fmt.Printf("Cache length: %d\n", len(cache))

	for inputKey, inputValue := range input {
		out, existsInCache := cache[inputKey]

		// If the value exists in the cache, use it
		if existsInCache {
			for _, cacheVal := range out {
				for range inputValue {
					output = addToMap(output, cacheVal)
				}
			}
			continue
		}

		if inputKey == 0 {
			for range inputValue {
				output = addToMap(output, 1)
			}
			cache[inputKey] = []int{1}
		} else if len(fmt.Sprintf("%d", inputKey))%2 == 0 {
			str := fmt.Sprintf("%d", inputKey)
			left := str[0 : len(str)/2]
			right := str[len(str)/2:]

			leftVal, _ := strconv.Atoi(left)
			rightVal, _ := strconv.Atoi(right)

			for range inputValue {
				output = addToMap(output, leftVal)
				output = addToMap(output, rightVal)
			}
			cache[inputKey] = []int{leftVal, rightVal}
		} else {
			calc := inputKey * 2024
			for range inputValue {
				output = addToMap(output, calc)
			}
			cache[inputKey] = []int{calc}
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

	cache := make(map[string]int)

	final := 0

	// Do each number from the input individually
	// for _, val := range input {
	// 	valueInput := map[int]int{
	// 		val: 1,
	// 	}
	//
	// 	for i := range 75 {
	// 		valueLength := 0
	// 		valueInput = blink(valueInput, cache)
	//
	// 		for _, val := range valueInput {
	// 			valueLength += val
	// 		}
	//
	// 		fmt.Printf("Input %d length: %d\n", i, valueLength)
	// 	}
	// 	for _, val := range valueInput {
	// 		final += val
	// 	}
	// }

	for _, val := range input {
		fmt.Println(val)
		final += blinkRecursive(val, 75, cache)
	}
	fmt.Printf("Length after 75 runs: %d", final)
}
