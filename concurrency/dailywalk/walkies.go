package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
Every morning, Alice and Bob go for a walk, and being creatures of habit,
they follow the same routine every day.

First, they both prepare,
grabbing sunglasses, perhaps a belt,
closing open windows, turning off ceiling fans, and pocketing their phones and keys.

Once they’re both ready, which typically takes each of them
between 60 and 90 seconds, they arm the alarm, which has a 60 second delay.

While the alarm is counting down, they both put on their shoes,
a process which tends to take each of them between 35 and 45 seconds.

Then they leave the house together and lock the door,
before the alarm has finished its countdown.

Write a program to simulate Alice and Bob’s morning routine.

Here’s some sample output from running a solution to this problem.
*/
type person struct {
	name string
}

const (
	readyMin      = 60
	readyMax      = 90
	alarmDuration = 60
	shoesMin      = 35
	shoesMax      = 45
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	persons := []person{{"Alice"}, {"Bob"}}
	// start up the persons channel
	gen := make(chan person, len(persons))
	// put persons on gen channel - will not block due to buffered channel
	for _, p := range persons {
		gen <- p
	}
	close(gen)

	// wg for persons to finish action every time
	var personwg sync.WaitGroup

	// the output channel for ready stage
	ready := make(chan person, len(persons))

	// ready stage
	personwg.Add(len(persons))
	// range the persons since channel is closed
	for p := range gen {
		go getReady(p, ready, &personwg)
	}
	personwg.Wait()
	close(ready)

	//time to arm the alarm once everyone is ready
	var alarmwg sync.WaitGroup
	alarmwg.Add(1)
	go alarm(&alarmwg)

	// alarm is running time to put shoes on
	// shoes stage
	personwg.Add(len(persons))
	for p := range ready {
		go putShoes(p, &personwg)
	}
	personwg.Wait()

	fmt.Println("Exiting and locking the door.")

	// alarm finishes at the very end
	alarmwg.Wait()
}

func getReady(p person, out chan<- person, wg *sync.WaitGroup) {
	defer wg.Done()
	d := getDuration(readyMin, readyMax)
	fmt.Printf("%s has started getting ready.\n", p.name)
	time.Sleep(d) // simulate duration
	fmt.Printf("%s spent %ds getting ready.\n", p.name, d/time.Second)
	out <- p // output person when finished
}

func putShoes(p person, wg *sync.WaitGroup) {
	defer wg.Done()
	d := getDuration(shoesMin, shoesMax)
	fmt.Printf("%s has started putting shoes on.\n", p.name)
	time.Sleep(d) // simulate duration
	fmt.Printf("%s spent %ds putting shoes on.\n", p.name, d/time.Second)
}

func alarm(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Arming the alarm.\n")
	time.Sleep(alarmDuration * time.Second) // simulate duration
	fmt.Printf("Alarm is armed.\n")
}

func getDuration(min int, max int) time.Duration {
	return time.Duration(rand.Intn(max-min)+min) * time.Second
}
