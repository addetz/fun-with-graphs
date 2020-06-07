package main

import utils "github.com/fun-with-graphs/daily-coding-problem"

func main() {
    m1 := [][]int {{1,2}, {3,4}}
    m2 := [][]int {{5,6}, {7,8}}
    prod := multiply(m1,m2,2)
    utils.PrintIntMatrix("M1", m1)
    utils.PrintIntMatrix("M2", m2)
    utils.PrintIntMatrix("Prod", prod)
}

func multiply(m1 [][]int, m2 [][]int, count int) [][]int{
    prod := make([][]int, count)
    for i := 0; i < count; i++ {
        prod[i] = make([]int, count)
    }
    for l := 0; l < count; l++ {
        for m := 0; m < count; m++ {
            prodSum := 0
            for n := 0; n < count; n++ {
                prodSum = prodSum + m1[l][n] * m2[n][m]
            }
            prod[l][m] = prodSum
        }
    }
    return prod
}
