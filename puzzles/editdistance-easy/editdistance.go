package main

import (
	"fmt"
	"strings"
)

/**
The edit distance between two strings refers to the minimum number of character
insertions, deletions, and substitutions required to change one string to the
other.
For example, the edit distance between “kitten” and “sitting” is three:
substitute the “k” for “s”, substitute the “e” for “i”, and append a “g”.

Given two strings, compute the edit distance between them.
*/
func main() {
	source := "kitten"
	target := "sitting"
	fmt.Printf("Edit distance between %s and %s is %d.\n",
		source, target, distance(source, target))

	source2 := "malamute"
	target2 := "salem"
	fmt.Printf("Edit distance between %s and %s is %d.\n",
		source2, target2, distance(source2, target2))
}

func distance(source string, target string) int {
	var difference int
	sources := strings.Split(source, "")
	targets := strings.Split(target, "")
	if len(targets) > len(sources) {
		difference = len(targets) - len(sources) // these are append diffs
		targets = targets[0:len(sources)]        // go just to the length of the target
	}
	if len(targets) < len(sources) {
		difference = len(sources) - len(targets) // these are remove diffs
		sources = sources[0:len(targets)]
	}
	for i := 0; i < len(targets); i++ {
		if sources[i] != targets[i] {
			difference++
		}
	}
	return difference
}
