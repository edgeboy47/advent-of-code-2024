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

// func compactFiles(block []string) []string {
// 	res := block
// 	rightPtr := len(block) - 1
// 	blanks := []struct {
// 		position, size int
// 	}{}
//
// 	return res
// }

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

func compactDiskmap(input string) int {

	type Block struct {
		position, size int
	}
	total := 0
	id := 0
	pos := 0
	isFile := true
	files := make(map[int]Block)
	blanks := []Block{}

	for _, val := range input {
		numBlocks, _ := strconv.Atoi(string(val))

		if isFile {
			files[id] = Block{position: pos, size: numBlocks}
			id += 1
		} else {
			blanks = append(blanks, Block{position: pos, size: numBlocks})
		}

		isFile = !isFile
		pos += numBlocks
	}

	fileIndex := id - 1

	// Look at each file, and look for a blank space where it can fit
	for fileIndex >= 0 {
		file := files[fileIndex]
		fileSize := file.size
		if fileSize == 0 {
			continue
		}

		for i := 0; i < len(blanks); i++ {
			blank := blanks[i]
			if blank.size == 0 {
				continue
			}
			// If the file can fit in the blank space
			if blank.size >= fileSize && file.position > blank.position {
				// put the file into the blank space
				newFilePos := blank.position
				file.position = newFilePos

				// put a new blank space where the file was
				// blanks = append(blanks, Block{position: file.position, size: fileSize})

				// calculate the size of the blank space after putting in the file
				remainingBlankSize := blank.size - fileSize
				blank.size = remainingBlankSize

				// If there is remaining space in the blank
				// that creates a new blank here and one where the file was
				if remainingBlankSize > 0 {
					blank.position = blank.position + fileSize
				}
				blanks[i] = Block{blank.position, blank.size}
			}
		}
		files[fileIndex] = file
		fileIndex--
	}

	// Create resulting fileblock string
	for fileIndex <= id {
		file := files[fileIndex]
		for j := range file.size {
			total += (file.position + j) * fileIndex
		}
		fileIndex++
	}

	return total
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
		defragChecksum := compactDiskmap(line)

		fmt.Printf("Checksum: %d\n", checksum)
		fmt.Printf("Defragmented Checksum: %d\n", defragChecksum)
	}
}
