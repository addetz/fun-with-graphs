package main

import (
    "errors"
    "fmt"
    "sync"
)
type machine struct {
    items map[int]item
    wg *sync.WaitGroup
    in <-chan selection // receive selection channel
    out chan<- item // send item channel
    err chan<- error // send error channel
    done <-chan struct{}  // receive done channel
}
type item struct {
    name string
    price int
}

func newMachine(wg *sync.WaitGroup, in <-chan selection, out chan<- item, err chan<- error, done <-chan struct{}) *machine{
    m := &machine{
        items: make(map[int]item),
        wg:    wg,
        in:    in,
        out:   out,
        err : err,
        done:  done,
    }
    m.items[1]= item {
        name:  "bounty",
        price: 3,
    }
    m.items[2] = item{
        name:  "snickers",
        price: 5,
    }
    return m
}

func (m *machine) start() {
    fmt.Println("Machine has started...")
    defer m.wg.Done()
    for {
        select {
        case <-m.done:
            fmt.Println("Machine has shutdown...")
            return
        case s := <-m.in :
            item, err := m.getItem(s)
            if err != nil {
                m.err <- err
                break
            }
            m.out <- item
        }
    }
}

func (m machine) getItem(s selection) (item, error) {
    i, ok := m.items[s.id]
    if !ok {
        return item{}, errors.New("item does not exist on machine")
    }
    totalSum := 0
    for _, c := range s.coins {
        totalSum = totalSum + c
    }
    if i.price > totalSum {
        return item{}, errors.New("not enough money supplied for selection")
    }
    return i, nil
}