package main

import (
    "fmt"
    "sort"
)

type thing struct{
    name string
    age int
}

type byFactor func(thing1 thing, thing2 thing) bool
func (bf byFactor) Sort(things []thing) {
    ts := &thingSorter{
        things:   things,
        byFactor: bf,
    }
    sort.Sort(ts)
}
type thingSorter struct {
    things []thing
    byFactor func(thing1 thing, thing2 thing) bool
}

func (t thingSorter) Len() int {
    return len(t.things)
}

func (t thingSorter) Less(i, j int) bool {
    return t.byFactor(t.things[i], t.things[j])
}

func (t thingSorter) Swap(i, j int) {
    t.things[i], t.things[j] = t.things[j], t.things[i]
}

func main() {
    things := []thing{
        {"adelina", 31},
        {"stuzzle", 34},
        {"alicia", 24},
    }
    printThings("Original", things)
    byAge := func(t1 thing, t2 thing) bool {
        return t1.age < t2.age
    }
    byFactor(byAge).Sort(things)
    printThings("Ordered", things)
}

func printThings(title string, things []thing) {
    fmt.Println(title)
    for i, v := range things {
        fmt.Printf("[%d]:%v \n", i, v)
    }
}
