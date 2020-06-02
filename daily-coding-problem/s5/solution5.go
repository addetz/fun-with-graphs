package main
import "C"
import "fmt"

/**
Given an array of integers, find the first missing positive integer
in linear time and constant space. In other words,
find the lowest positive integer that does not exist in the array.
The array can contain duplicates and negative numbers as well.

For example, the input [3, 4, -1, 1] should give 2.
The input [1, 2, 0] should give 3.

You can modify the input array in-place.
 */
func main() {
    input := []int {3,4,-1,1}

    // Segregate positive numbers from others i.e., move all non-positive numbers to left side.
    // Now we can ignore non-positive elements and consider only the part of array
    // which contains all positive elements.
    neg := segregate(input)
    // We traverse the array containing all positive numbers and
    // to mark presence of an element x, we change the sign of value at index x to negative.
    positive := input[neg:]
    for i, v := range positive {
        positive[i] = v * -1
    }
    // We traverse the array again and print the first index which has positive value.
    // In the following code, findMissingPositive() function does this part.
    firstPositive := len(positive)
    for i, v := range positive{
        if v > 0 {
            firstPositive = i + 1
            break
        }
    }

    fmt.Printf("The missing positive integer is %d \n", firstPositive)
}

func segregate(input []int) int {
    neg := 0
    for i, v := range input {
        if v < 0 {
            input[i] = input[neg]
            input[neg] = v
            // increment count of non-positive
            // integers
            neg++
        }
    }
    return neg
}
