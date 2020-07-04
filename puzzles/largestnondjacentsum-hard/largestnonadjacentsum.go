package main

import (
	"fmt"
	"math"
)

/**
Given a list of integers, write a function that returns the
largest sum of non-adjacent numbers. Numbers can be 0 or negative.

For example,
[2, 4, 6, 2, 5] should return 13, since we pick 2, 6, and 5.
[5, 1, 1, 5]should return 10, since we pick 5 and 5.

Follow-up: Can you do this in O(N) time and constant space?
*/
func main() {
	input1 := []float64{2, 4, 6, 2, 5}
	input2 := []float64{5, 1, 1, 5}
	fmt.Printf("Largest sum of %v is %d.\n", input1, int(nonAdjSum(input1)))
	fmt.Printf("Largest sum of %v is %d.\n", input2, int(nonAdjSum(input2)))
}

func nonAdjSum(input []float64) float64 {
	var incl float64
	var excl float64
	for _, in := range input {
		temp := incl
		incl = math.Max(incl, excl+in)
		excl = temp
	}
	return math.Max(incl, excl)
}
