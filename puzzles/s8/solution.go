package main

import (
    "fmt"
)

/***
Smallest missing positive integer
 */
func main() {
    allNeg := []int{-1,-3}
    allPos := []int{1,2,3}
    allPos2 := []int{1, 3, 6, 4, 1, 2}
    fmt.Printf("All neg: %d \n", findMissing(allNeg))
    fmt.Printf("All Pos: %d \n", findMissing(allPos))
    fmt.Printf("All Pos2: %d \n", findMissing(allPos2))
}

func findMissing(input []int) int {
    m := make(map[int]int)
    for _, val := range input {
        if val > 0 {
            m[val] = val
        }
    }
    positive := 1
    for {
        if _, ok := m[positive]; !ok {
            break
        } else {
            positive++
        }
    }
    return positive
}


