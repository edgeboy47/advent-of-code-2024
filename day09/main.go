package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func diskmapToFileBlocks(disk string) []string {
	res := []string{}
	id := 0
	isFile := true

	for _, val := range disk {
		numBlocks, _ := strconv.Atoi(string(val))

		for range numBlocks {
			if isFile {
				res = append(res, fmt.Sprintf("%d", id))
			} else {
				res = append(res, ".")
			}
		}

		if isFile {
			id += 1
		}
		isFile = !isFile
	}

	return res
}

func compactFileBlock(block []string) []string {
  leftPtr := 0
  rightPtr := len(block) - 1
  fileBlocks := block


  for leftPtr < rightPtr {
    if fileBlocks[rightPtr] == "." {
      rightPtr -= 1
      continue
    }

    if fileBlocks[leftPtr] == "." && fileBlocks[rightPtr] != "." {
      temp := fileBlocks[leftPtr]
      fileBlocks[leftPtr] = fileBlocks[rightPtr]
      fileBlocks[rightPtr] = temp
      rightPtr -= 1
    }

    leftPtr++
  }

	return fileBlocks
}

func calculateChecksum(block []string) int {
  res := 0

  for i, val := range block {
    if val == "." {
      continue
    }

    num, _ := strconv.Atoi(val)
    res += i * num
  }

  return res
}

func main() {
	inputFile := "input.txt"
	file, err := os.Open(inputFile)

	if err != nil {
		println("err reading file", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
    fileBlocks := diskmapToFileBlocks(line)
    compactBlock := compactFileBlock(fileBlocks)
    checksum := calculateChecksum(compactBlock)

    fmt.Printf("Checksum: %d\n", checksum)
  }
}
