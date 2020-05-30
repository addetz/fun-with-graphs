package main

import (
    "fmt"
)

/**
There's a staircase with N steps, and you can climb 1 or 2 steps at a time.
Given N, write a function that returns the number of unique ways you can climb the staircase.
The order of the steps matters.

For example, if N is 4, then there are 5 unique ways:

1, 1, 1, 1
2, 1, 1
1, 2, 1
1, 1, 2
2, 2

What if, instead of being able to climb 1 or 2 steps at a time, you could climb
any number from a set of positive integers X?
For example, if X = {1, 3, 5}, you could climb 1, 3, or 5 steps at a time.
Generalize your function to take in X.
 */
type climb struct {
    total int
    path string
    isIncomplete bool
}
func main() {
    n := 5
    steps := []int{1, 2, 3}
    climbs := computeClimbs(n, steps)
    fmt.Printf("Climbs count: %d \n", len(climbs))
    for _, c := range climbs {
        fmt.Printf("%s \n", c.path)
    }
}

func computeClimbs(n int, steps []int) []*climb {
    climbs := make([]*climb, 0)
    for _, s := range steps {
        if s <= n {
            climbs = append(climbs, &climb {
                    total: s,
                    path: fmt.Sprintf("%d", s),
                    isIncomplete: s < n,
            })
        }
    }

    cont := len(climbs) > 0
    for cont {
        climbs, cont = computeClimbsHelper(climbs, n, steps)
    }
    return climbs
}

func computeClimbsHelper(prev []*climb, target int, steps []int) ([]*climb, bool) {
    updatedClimbs := make([]*climb,0)
    found := false
    for _, cl := range prev {
        if cl.isIncomplete {
            found = true
            updatedClimbs = append(updatedClimbs, update(cl, target, steps)...)
            continue
        }
        updatedClimbs = append(updatedClimbs, cl)
    }
    return updatedClimbs, found
}

func update(cl *climb, target int, steps []int) []*climb {
    uc := make([]*climb,0)
    for _, s := range steps {
        if cl.total + s <= target {
            newC := &climb{
                total: cl.total + s,
                isIncomplete: cl.total + s < target,
                path:  fmt.Sprintf("%s,%d", cl.path, s),
            }
            uc = append(uc, newC)
        }
    }
    return uc
}