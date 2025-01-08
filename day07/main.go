package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type TreeNode struct {
	leftNode   *TreeNode
	middleNode *TreeNode
	rightNode  *TreeNode
	val        int
}

func addNode(root *TreeNode, val int) {
	if root == nil || reflect.DeepEqual(*root, TreeNode{}) {
		newNode := TreeNode{
			leftNode:   nil,
			middleNode: nil,
			rightNode:  nil,
			val:        val,
		}
		*root = newNode
		return
	}

	if root.leftNode != nil && root.middleNode != nil && root.rightNode != nil {
		addNode(root.leftNode, val)
		addNode(root.middleNode, val)
		addNode(root.rightNode, val)
		return
	}

	concatVal, _ := strconv.Atoi(fmt.Sprintf("%d%d", root.val, val))
	newLeftNode := TreeNode{
		leftNode:   nil,
		middleNode: nil,
		rightNode:  nil,
		val:        root.val + val,
	}

	newMiddleNode := TreeNode{
		leftNode:   nil,
		middleNode: nil,
		rightNode:  nil,
		val:        concatVal,
	}

	newRightNode := TreeNode{
		leftNode:   nil,
		middleNode: nil,
		rightNode:  nil,
		val:        root.val * val,
	}

	root.leftNode = &newLeftNode
	root.middleNode = &newMiddleNode
	root.rightNode = &newRightNode

	return
}

func search(root *TreeNode, val int) bool {
	if root == nil {
		return false
	}

	if root.val == val {
		return true
	}

	return search(root.leftNode, val) || search(root.middleNode, val) || search(root.rightNode, val)
}

func main() {
	inputFile := "input.txt"
	file, err := os.Open(inputFile)

	if err != nil {
		println("err reading file", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	total := 0

	for scanner.Scan() {
		line := scanner.Text()
		root := TreeNode{}

		regex, _ := regexp.Compile(`(\d+)\: (.*)`)
		matches := regex.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			value, _ := strconv.Atoi(match[1])
			vals := strings.Split(match[2], " ")

			for _, val := range vals {
				num, _ := strconv.Atoi(val)
				addNode(&root, num)
			}

			if search(&root, value) {
				total += value
			}
		}
	}

	fmt.Printf("Total calibration result: %d\n", total)
}
