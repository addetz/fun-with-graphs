package main

import (
	"fmt"
	"sort"
)

/**
N trees in a forest, k'th tree is at location x[k], y[k]
Build the widest vertical path
Path must be between 2 trees and cannot be infinite
Find widest possible path, a path of at least 1 will be possible
*/
func main() {
	x1 := []int{1, 8, 7, 3, 4, 1, 8}
	y1 := []int{6, 4, 1, 8, 5, 1, 7}
	x2 := []int{5, 5, 5, 7, 7, 7}
	y2 := []int{3, 4, 5, 1, 3, 7}
	x3 := []int{4, 1, 5, 4}
	y3 := []int{4, 5, 1, 3}
	fmt.Printf("Widest path size is %d \n", widestPath(x1, y1))
	fmt.Printf("Widest path size is %d \n", widestPath(x2, y2))
	fmt.Printf("Widest path size is %d \n", widestPath(x3, y3))
	fmt.Printf("Widest path size is %d \n", widestPath2(x1, y1))
	fmt.Printf("Widest path size is %d \n", widestPath2(x2, y2))
	fmt.Printf("Widest path size is %d \n", widestPath2(x3, y3))
}
func widestPath2(X []int, Y []int) int {
	if len(X) < 2 {
		return 0
	}
	// invalid input, as Y must be same length as X
	if len(X) != len(Y) {
		return 0
	}
	// sort the X's ascending
	sort.Ints(X)
	// we start at the second tree and find the diff with previous tree
	max := 0
	for i := 1; i < len(X); i++ {
		// find the max of consecutive trees
		diff := X[i] - X[i-1]
		if diff > max {
			max = diff
		}
	}

	return max
}

func widestPath(x1 []int, _ []int) int {
	// need at least 2 trees for the path to not be infinite
	if len(x1) < 1 {
		return 0
	}
	// sort the trees by x coordinate
	sort.Ints(x1)
	// if there are only 2 trees we simply return their difference
	if len(x1) == 2 {
		return x1[1] - x1[0]
	}
	// we start at the second tree and find the diff with previous tree
	max := 0
	for i := 1; i < len(x1); i++ {
		// find the max of consecutive trees
		diff := x1[i] - x1[i-1]
		if diff > max {
			max = diff
		}
	}
	return max
}
