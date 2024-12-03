package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
  file, err := os.Open("input.txt")

  diff := 0.0
  leftList, rightList := []int{}, []int{}
  if (err != nil) {
    println("err reading file", err)
  }
  defer file.Close()
  
  scanner := bufio.NewScanner(file)

  // Part 1
  for scanner.Scan() {
    line := scanner.Text()
    vals := strings.Split(line, "   ")
    line1, line2 := vals[0], vals[1]

    val1, err1 := strconv.Atoi(line1)
    val2, err2 := strconv.Atoi(line2)

    if (err1 != nil || err2 != nil) {
    }
    leftList = append(leftList, val1)
    rightList = append(rightList, val2)
  }

  sort.Ints(leftList)
  sort.Ints(rightList)

  for i := 0; i < len(leftList); i++ {
    diff += math.Abs(float64(leftList[i] - rightList[i]))
  }

  fmt.Printf("diff: %g\n", diff)

  // Part 2
  rightListFrequency := make(map[int]int)
  // Create frequency table for each number in list 2
  for _, i := range rightList {
    count, exists := rightListFrequency[i]
    if (exists) {
      rightListFrequency[i] = count + 1
    } else {
      rightListFrequency[i] = 1
    }
  }

  similarity := 0.0

  for _, i := range leftList {
    count, exists := rightListFrequency[i]
    if (exists) {
      similarity += float64(i * count)
    }
  }

  fmt.Printf("similarity score: %g\n", similarity)
}
