package main

import (
	"fmt"
	"strings"
)

/**
Given a string of round, curly, and square open and closing brackets,
return whether the brackets are balanced (well-formed).

For example, given the string "([])[]({})", you should return true.

Given the string "([)]" or "((()", you should return false.
*/
func main() {
	input1 := "([])[]({})"
	input2 := "([)]"
	input3 := "((()"
	fmt.Printf("%s has returned balaced: %t\n", input1, balanceBrackets(input1))
	fmt.Printf("%s has returned balaced: %t\n", input2, balanceBrackets(input2))
	fmt.Printf("%s has returned balaced: %t\n", input3, balanceBrackets(input3))
}

func balanceBrackets(input string) bool {
	openingBrackets := newStack()
	brackets := strings.Split(input, "")
	for _, b := range brackets {
		// we push opening brackets on their stack and don't care
		if isOpeningBracket(b) {
			openingBrackets.push(b)
			continue
		}
		if isClosingBracket(b) {
			openB, valid := openingBrackets.pop()
			if !valid {
				// we encountered a close bracket with no openings on  the stack
				return false
			}
			if !matchBrackets(openB, b) {
				return false
			}
		}
		// if it isn't a bracket we throw it away and continue
		continue
	}
	//all done with the input - the opening stack should be empty now
	return openingBrackets.isEmpty()
}

func isOpeningBracket(b string) bool {
	return b == "(" || b == "[" || b == "{"
}

func isClosingBracket(b string) bool {
	return b == ")" || b == "]" || b == "}"
}

func matchBrackets(b1, b2 string) bool {
	switch b1 {
	case "(":
		return b2 == ")"
	case "[":
		return b2 == "]"
	case "{":
		return b2 == "}"
	default:
		return false
	}
}
