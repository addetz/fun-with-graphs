package main

import "fmt"

/**
Identify the submatrices of 2x2 in a 3x3 matrix.
What is the complexity of the algorithm that you have used?
 */
type submatrix struct {
    vals  [][]int
    start []int
}
func (s *submatrix) print() {
    fmt.Printf("[%d %d]: \n", s.start[0], s.start[1])
    for _,vals := range s.vals {
        for _, v := range vals {
            fmt.Printf("%d ", v)
        }
        fmt.Printf("\n")
    }
    fmt.Printf("\n")
}
func main() {
    input := [][]int{{1,2,3,4}, {4,5,6,7}, {7,8,9,10}}
    // size of matrix we are looking for
    k := 2
    sm := calculateSubmatrices(input, k)
    for _, s := range sm {
        s.print()
    }
}

func calculateSubmatrices(input [][]int, k int) []*submatrix{
    var sm []*submatrix
    for i := 0; i <= len(input) - k; i++ {
        for j := 0; j <= len(input[i]) - k; j++ {
            sm = append(sm, calculateSubmatrix(input, k, i, j))
        }
    }
    return sm
}

func calculateSubmatrix(input [][]int, k int, x int, y int) *submatrix {
    vals := make([][]int, k)
    for i := 0; i < k ; i ++ { // row
        row := make([]int, k)
        for j := 0; j < k ; j ++ { //col
            row[j] = input[x+i][y+j]
        }
        vals[i] = row
    }

    return &submatrix{
        vals:  vals,
        start: []int{x,y},
    }
}
