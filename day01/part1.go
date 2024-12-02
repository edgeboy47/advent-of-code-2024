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
  list1, list2 := []int{}, []int{}
  if (err != nil) {
    println("err reading file", err)
  }
  defer file.Close()
  
  scanner := bufio.NewScanner(file)

  for scanner.Scan() {
    line := scanner.Text()
    vals := strings.Split(line, "   ")
    line1, line2 := vals[0], vals[1]

    val1, err1 := strconv.Atoi(line1)
    val2, err2 := strconv.Atoi(line2)

    if (err1 != nil || err2 != nil) {
    }
    list1 = append(list1, val1)
    list2 = append(list2, val2)
  }

  sort.Ints(list1)
  sort.Ints(list2)

  for i := 0; i < len(list1); i++ {
    diff += math.Abs(float64(list1[i] - list2[i]))
  }

  fmt.Printf("diff: %g", diff)
}
