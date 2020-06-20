package main

import (
    "fmt"
    "sync"
)

// vending machine with no change but working on concurrent routines
func main() {
    coins := []int{1,2,5}

    //start the synch channels
    var customersWg sync.WaitGroup
    var vendingWg sync.WaitGroup
    customers := make(chan selection)
    soldItems := make(chan item)
    err := make(chan error)
    done := make(chan struct{})

    //start up the vending machine
    machine := newMachine(&vendingWg, customers, soldItems, err, done)
    go machine.start()

    //start some customers routines and wait for them to finish in main
    adelina := newGopher("Adelina", selection{1, coins[:2]}, &customersWg, customers, soldItems, err)
    stuzzle := newGopher("Stuzzle" ,selection{2, coins[2:]}, &customersWg, customers, soldItems, err)
    broke := newGopher("Broke" ,selection{2, coins[:1]}, &customersWg, customers, soldItems, err)
    go adelina.start()
    go stuzzle.start()
    go broke.start()
    customersWg.Wait()

    // close the vending machine once the customers have left
    close(done)
    vendingWg.Wait()
    fmt.Println("Main has shutdown. Goodbye!")
}
