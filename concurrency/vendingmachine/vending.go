package main

import (
    "fmt"
    "sync"
)

/*
Design a vending machine that will take coins as input and chocolate as output.

Let's imagine that Gophers have coins of value 1, 2 and 5.
Gopher can choose the item and then has to insert money and
wait for the vending machine to dispense the item.
 */
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
