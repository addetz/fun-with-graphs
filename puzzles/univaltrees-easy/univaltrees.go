package main

import "fmt"

/*
A unival tree (which stands for "universal value") is a tree where all nodes under it have the same value.

Given the root to a binary tree, count the number of unival subtrees.
*/
type node struct {
	value       int
	name        string
	left, right *node
}

func main() {
	root := setupTree()
	var count int
	countUnival(root, &count)
	fmt.Printf("Unival: %d \n", count)
}

func countUnival(node *node, count *int) bool {
	// this could happen for a lopsided node
	if node == nil {
		return true
	}
	// we are a leaf
	if node.left == nil && node.right == nil {
		*count++
		return true
	}
	// we are not a leaf
	isUniL := countUnival(node.left, count)
	isUniR := countUnival(node.right, count)
	if isUniL && isUniR {
		if isUnivalNode(node) {
			*count++
			return true
		}
		return false
	}
	return false
}

func isUnivalNode(node *node) bool {
	leftEq := true
	rightEq := true
	if node.left != nil && node.left.value != node.value {
		leftEq = false
	}
	if node.right != nil && node.right.value != node.value {
		rightEq = false
	}
	return leftEq && rightEq
}

func setupTree() *node {
	/*
			           root - 0
			       l1-0         r1 - 1
			   l1l2-0   l1r2-0      r1r2-1
		    Expected unival count - 5
	*/
	r1r2 := &node{
		value: 1,
		name:  "r1r2",
	}
	l1l2 := &node{
		value: 0,
		name:  "l1l2",
	}
	l1r2 := &node{
		value: 0,
		name:  "l1r2",
	}
	l1 := &node{
		value: 0,
		name:  "l1",
		left:  l1l2,
		right: l1r2,
	}
	r1 := &node{
		value: 1,
		name:  "r1",
		right: r1r2,
	}
	return &node{
		value: 0,
		name:  "root",
		left:  l1,
		right: r1,
	}
}
