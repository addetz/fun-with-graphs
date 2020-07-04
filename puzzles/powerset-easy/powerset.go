package main

import (
	"fmt"
	"math"
	"strings"
)

/**
The power set of a set is the set of all its subsets.
Write a function that, given a set, generates its power set.

For example, given the set {1, 2, 3},
it should return {{}, {1}, {2}, {3}, {1, 2}, {1, 3}, {2, 3}, {1, 2, 3}}.

You may also use a list or array to represent a set.
*/
func main() {
	input := []string{"a", "b", "c", "d", "e"}
	fmt.Printf("Powerset of %v is %v\n", input, powerset(input))
}

func powerset(input []string) []string {
	powerSetSize := int(math.Pow(2, float64(len(input))))
	setSize := len(input)
	pSet := make([]string, powerSetSize)
	for count := 0; count < powerSetSize; count++ {
		if count == 0 {
			pSet[count] = "E"
			continue
		}
		var value strings.Builder
		for j := 0; j < setSize; j++ {
			// Check if jth bit in the counter is set
			b := count & (1 << j)
			if b > 0 {
				value.WriteString(input[j])
			}
		}
		pSet[count] = value.String()
	}
	return pSet
}
