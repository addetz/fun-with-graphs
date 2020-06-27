package main

import "fmt"

/**
Given a list of numbers and a number k, return whether any two numbers from the list add up to k.

For example, given [10, 15, 3, 7] and k of 17, return true since 10 + 7 is 17.

Bonus: Can you do this in one pass?
*/
func main() {
	input := []int{10, 15, 14, 3, 7}
	target := 17
	coeff, found := findMySum(input, target)
	fmt.Printf("Coeff: [%d, %d] \n", coeff[0], coeff[1])
	fmt.Printf("Found is %t \n", found)

}

func findMySum(input []int, target int) ([]int, bool) {
	remainders := make(map[int]int, 0)
	for _, v := range input {
		first, ok := remainders[v]
		if ok {
			return []int{first, v}, true
		}
		remainders[target-v] = v
	}
	return []int{0, 0}, false
}
