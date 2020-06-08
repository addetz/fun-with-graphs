package utils

import "fmt"

func PrintInt(text string, output []int) {
    fmt.Printf("%s: \n", text)
    for _, v := range output {
        fmt.Printf(" %d ", v)
    }
    fmt.Printf("\n")
}

func PrintIntMatrix(text string, output [][]int) {
    fmt.Printf("%s: \n", text)
    for _, r := range  output {
        fmt.Printf("[")
        for _, v := range r {
            fmt.Printf(" %d ", v)
        }
        fmt.Printf("]\n")
    }
}