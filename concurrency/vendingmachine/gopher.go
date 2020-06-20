package main

import (
    "fmt"
    "math/rand"
    "sync"
    "time"
)
type gopher struct {
    name string
    selection selection
    wg *sync.WaitGroup
    out chan<- selection // send only selection channel
    in  <-chan item // receive only item channel
    err <-chan error // receive only error channel
}
type selection struct {
    id int
    coins []int
}
func newGopher(name string, selection selection, wg *sync.WaitGroup, out chan<- selection, in <-chan item, err <-chan error) *gopher {
    wg.Add(1)
    return &gopher{
        name: name,
        selection: selection,
        wg: wg,
        out: out,
        in: in,
        err: err,
    }
}

func (g gopher) purchase() {
    // sleep for a little bit while the gopher thinks about what they want to buy
    n := rand.Intn( 6 ) * 100
    time.Sleep(time.Duration(n) * time.Millisecond)
    fmt.Printf("Gopher %s purchases %v\n", g.name, g.selection)
    g.out <- g.selection
}

func (g *gopher) start() {
    fmt.Printf("Gopher %s has started... \n", g.name)
    // attempt to purchase at some random time
    defer g.wg.Done()
    g.purchase()
    for {
        select {
        case item := <-g.in:
            fmt.Printf("Gopher %s receives %v\n", g.name, item)
            fmt.Printf("Gopher %s closed.\n", g.name)
            return
        case err := <-g.err:
            fmt.Printf("Gopher %s receives err %v\n", g.name, err)
            fmt.Printf("Gopher %s closed.\n", g.name)
            return
        }
    }
}