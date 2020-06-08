package main

import "fmt"

func main() {
    fmt.Println("Series: ", fib_series(8))
    fmt.Println("Rec: ", fib_rec(8))
}

// Iterative method
func fib_series(n int) int {
    f := make([]int, n+1, n+2)
    if n < 2 {
        f = f[0:2]
    }
    f[0] = 0
    f[1] = 1
    var i int
    for i = 2; i <= n; i++ {
        f[i] = f[i-1] + f[i-2]
    }
    return f[n]
}

// Recursive method
func fib_rec(n int) int {
    if n <= 1 {
        return n
    }
    return fib_rec(n-1) + fib_rec(n-2)
}