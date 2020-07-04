package main

import (
	"fmt"
	"strconv"
)

/**
Given the mapping a = 1, b = 2, ... z = 26,
and an encoded message, count the number of ways it can be decoded.

For example, the message '111' would give 3,
since it could be decoded as 'aaa', 'ka', and 'ak'.

You can assume that the messages are decodable.
*/
const maxLetterNumber = 26

func main() {
	input1 := "123" //{1,2,3}, {12,3}, {1,23}
	input2 := "127" //{1,2,7}, {12,7}
	fmt.Printf("Decodable ways %s:%d \n", input1, decode(input1))
	fmt.Printf("Decodable ways %s:%d \n", input2, decode(input2))
}

func decode(input string) int {
	// one single char can always be decoded 1 way unless, it's 0
	if len(input) == 1 && canDecode(input) {
		return 1
	}
	// 2 chars can be decoded 3 ways
	if len(input) == 2 {
		if canDecode(input) {
			return 2
		}
		return 1
	}
	// decompose a longer char
	return decode(input[:1])*decode(input[1:]) +
		decode(input[:2])*decode(input[2:]) - 1
}

func canDecode(s string) bool {
	// if we get a 01 break then it's not decodable
	if s[0] == '0' {
		return false
	}
	num, err := strconv.Atoi(s)
	if err != nil { // can't convert to number
		return false
	}
	return num <= 26
}
