package main

import "fmt"

type fraction struct {
    num, denom int
}

func (frac fraction) String() string {
    return fmt.Sprintf("%d/%d", frac.num, frac.denom)
}
/**
A Farey sequence consists of reduced fractions with values between zero and one.
The denominators of the fractions are less than or equal to m,
and organized in ascending order. This sequence is called a Farey series.
*/
func main() {
    // main method
    for num := 1; num <= 3; num++ {
        l := fraction{0, 1}
        r := fraction{1, 1}
        fmt.Printf("F(%d): %s ", num, l)
        farey(l, r, num)
        fmt.Println(r)
    }
}

// farey method
func farey(l fraction, r fraction, num int) {
    frac := fraction{l.num + r.num, l.denom + r.denom}
    if frac.denom <= num {
        farey(l, frac, num)
        fmt.Print(frac, " ")
        farey(frac, r, num)
    }
}