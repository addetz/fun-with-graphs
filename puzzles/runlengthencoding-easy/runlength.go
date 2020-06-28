package main

import (
	"fmt"
	"strconv"
	"strings"
)

/**
Run-length encoding is a fast and simple method of encoding strings.
The basic idea is to represent repeated successive characters as a single count
and character.
For example, the string "AAAABBBCCDAA" would be encoded as "4A3B2C1D2A".

Implement run-length encoding and decoding.
You can assume the string to be encoded have no digits
and consists solely of alphabetic characters.
You can assume the string to be decoded is valid.
*/
func main() {
	input1 := "AAAABBBCCDAA"
	output1 := encode(input1)
	fmt.Printf("%s encodes to %s\n", input1, output1)
	decode1 := decode(output1)
	fmt.Printf("%s decodes to %s\n", output1, decode1)
	input2 := "4A10B3C1D"
	decode2 := decode(input2)
	fmt.Printf("%s decodes to %s\n", input2, decode2)
}

func encode(input string) string {
	var output strings.Builder
	inputs := strings.Split(input, "")
	var prev string
	var count int
	for _, c := range inputs {
		if prev != c {
			if count != 0 {
				output.WriteString(fmt.Sprintf("%d%s", count, prev))
			}
			prev = c
			count = 1
			continue
		}
		count++
	}
	// add the last count also
	output.WriteString(fmt.Sprintf("%d%s", count, prev))
	return output.String()
}

func decode(input string) string {
	var output strings.Builder
	inputs := strings.Split(input, "")
	var count int
	var decodeNext bool
	for _, c := range inputs {
		i, err := toNumeric(c)
		// the char is not a digit and
		if err != nil {
			// we didn't expect to decode next so dump the char in output
			if !decodeNext {
				output.WriteString(c)
				continue
			}
			// the decode next was true so we expand the count on the char
			for k := 0; k < count; k++ {
				output.WriteString(c)
			}
			//we've expanded, time to reset the decode
			decodeNext = false
			continue
		}
		//else we have a number, check
		// if a previous number was received and if so add it to the number
		if decodeNext {
			count = count*10 + i
			continue
		}
		count = i
		decodeNext = true
	}
	return output.String()
}

func toNumeric(s string) (int, error) {
	return strconv.Atoi(s)
}
