package main

import (
	"testing"
)

func TestAddNode(t *testing.T) {
	t.Run("add node correctly", func(t *testing.T) {
		root := TreeNode{}

		addNode(&root, 10)
		addNode(&root, 19)

		expectedLeft := 29
		actualLeft := root.leftNode.val

		expectedRight := 190
		actualRight := root.rightNode.val

		if expectedLeft != actualLeft || expectedRight != actualRight {
			t.Errorf("Error with adding node, expected %d and %d, got %d and %d", expectedLeft, expectedRight, actualLeft, actualRight)
		}
	})

	t.Run("Add deeply nested node", func(t *testing.T) {
		root := TreeNode{
			leftNode:   nil,
			middleNode: nil,
			rightNode:  nil,
			val:        11,
		}

		addNode(&root, 6)
		addNode(&root, 16)
		addNode(&root, 20)

		expected := 292
		actual := root.leftNode.rightNode.leftNode.val

		if expected != actual {
			t.Errorf("Error with adding deep nodes, expected %d got %d", expected, actual)
		}
	})
}

func TestSearch(t *testing.T) {
	t.Run("search deeply nested node", func(t *testing.T) {
		root := TreeNode{}
		addNode(&root, 11)
		addNode(&root, 6)
		addNode(&root, 16)
		addNode(&root, 20)

		expected := true
		actual := search(&root, 292)

		if expected != actual {
			t.Errorf("Error with search")
		}
	})
}
