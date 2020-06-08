package main

import utils "github.com/fun-with-graphs/puzzles"

/**
Given an array of integers, return a new array such that each element at index
i of the new array is the product of all the
numbers in the original array except the one at i.

For example, if our input was [1, 2, 3, 4, 5],
the expected output would be [120, 60, 40, 30, 24].
If our input was [3, 2, 1], the expected output would be [2, 3, 6].

Follow-up: what if you can't use division?
 */
func main() {
    input1 := []int{3,2,1}
    input2 := []int{1,2,3,4,5}
    utils.PrintInt("Input", input1)
    utils.PrintInt("Output", product(input1))
    utils.PrintInt("Input", input2)
    utils.PrintInt("Output", product(input2))
}

func product(input []int) []int {
    if len(input) == 1 {
        return []int{0}
    }
    prod := make([]int, len(input))
    temp := 1

    // traverse the array from the left to right -
    // put the product up to index i in the array
    for i, v := range input{
        prod[i] = temp
        temp = temp * v
    }

    temp = 1
    // traverse the array from the right to left and
    // put the product of everything up to that number in the array
    for i := len(input)-1 ; i >= 0 ; i-- {
        prod[i] = prod[i] * temp
        temp = temp * input[i]
    }
    return prod
}
