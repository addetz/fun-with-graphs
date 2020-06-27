package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

/**
Find the max of a group of siblings.
Given  a number, find it's maximum  sibling from the family.
*/
func main() {
	input1 := 319
	fmt.Printf("Max sibling of %d is %d \n", input1, maxSibling(input1))
	fmt.Printf("Max sibling of %d is %d \n", input1, maxSibling2(input1))
	input2 := 30109
	fmt.Printf("Max sibling of %d is %d \n", input2, maxSibling(input2))
	fmt.Printf("Max sibling of %d is %d \n", input2, maxSibling2(input2))
}

func maxSibling2(n int) int {
	// guarding against negative numbers as it will break the split
	if n < 0 {
		return 0
	}
	digits := strings.Split(fmt.Sprintf("%d", n), "")
	// n was a 1 digit number, return it as is
	if len(digits) == 1 {
		return n
	}
	//sorted asc digits - string sort will still work
	sort.Strings(digits)
	reverse := make([]string, len(digits))
	lastIndex := len(digits) - 1
	for i := lastIndex; i > -1; i-- {
		reverse[lastIndex-i] = digits[i]
	}
	// rejoin sorted slice
	rejoined := strings.Join(reverse, "")
	// parse the new max, return 0 in case of foul play
	max, err := strconv.Atoi(rejoined)
	if err != nil {
		return 0
	}
	return max
}

func maxSibling(n int) int {
	// we will stop the solution if the
	// input is negative as it breaks our other logic
	if n < 0 {
		return 0
	}
	digits := strings.Split(fmt.Sprintf("%d", n), "")
	//sort the digits descending order
	sort.Slice(digits, func(i, j int) bool {
		return digits[i] > digits[j]
	})
	// put the sorted list back together
	joined := strings.Join(digits, "")
	//disregard error as all numbers came from the input int
	max, _ := strconv.Atoi(joined)
	return max
}
