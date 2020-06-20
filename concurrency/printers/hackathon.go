package main

import (
	"fmt"
	"sync"
	"time"
)

/*
During busy hackaton 3D printers are heavily used.
Write a simulation of hackers trying to access 3D printers.

In the hackerspace, there are 3 3D printers. T
here are 7 hackers that are interested in using the printers.

If the hacker can't access the printer for more than 5 seconds,
he gets annoyed and quits the hackaton.
Hackers use printers for random interval from 1 to 10 seconds and
usually they need to use the printer at least twice,
because nothing is perfect for the first time.
*/
const (
	printerCount = 3
	hackerCount  = 7
	hackerEnding = 1 * time.Minute
)

var printers []*printer

func main() {
	printers = make([]*printer, printerCount)
	hackers := make([]*hacker, hackerCount)

	// buffered channel for jobs equal to the amount of printers we have
	jobs := make(chan job, printerCount)
	results := make(chan string, hackerCount)
	done := make(chan struct{})
	var printerwg sync.WaitGroup
	var hackerwg sync.WaitGroup
	defer close(jobs)
	defer close(results)

	for i := 0; i < printerCount; i++ {
		printers[i] = newPrinter(fmt.Sprintf("printer-%d", i),
			jobs, results, done, &printerwg)
	}
	printerwg.Add(printerCount)
	for _, p := range printers {
		go p.start()
	}
	for i := 0; i < hackerCount; i++ {
		hackers[i] = newHacker(fmt.Sprintf("hacker-%d", i),
			jobs, results, done, &hackerwg)
	}
	hackerwg.Add(hackerCount)
	for _, h := range hackers {
		go h.start()
	}
	printerwg.Add(1)
	// exit when the hackathon ends and close the printers
	go func(done chan struct{}) {
		defer printerwg.Done()
		timer := time.NewTimer(hackerEnding)
		<-timer.C
		fmt.Println("Hackathon is ended! Byeee!")
		close(done)
	}(done)

	hackerwg.Wait()
	printerwg.Wait()
	fmt.Println("Hackathon has ended! Goodbye.")
}
