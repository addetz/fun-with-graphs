package main

import (
	"fmt"
	"sync"
)

/*
A small internet café in a village just outside of Manilla has 8 computers,
which are available on a first-come first-serve basis.
When all the computers are taken,
the next person in line has to wait until a computer frees up.

This morning several groups of tourists, 25 people in all, are waiting when the doors open.

Each person spends between 15 minutes and 2 hours online.

Write a program to simulate the computer usage at the café.
*/
const (
	computerCount = 8
	userCount     = 25
)

type computer struct {
	name string
}

func main() {
	var userwg sync.WaitGroup
	computers := make(chan computer, computerCount)
	users := make(chan user, userCount)
	defer close(computers)

	// put computers in the channel as they are all available in the beginning
	for i := 0; i < computerCount; i++ {
		computers <- computer{fmt.Sprintf("computer-%d", i)}
	}
	// put the users in a channel as they are queued up for using a computer
	for i := 0; i < userCount; i++ {
		users <- newUser(fmt.Sprintf("user-%d", i), &userwg, computers)
	}
	close(users) // all the users have been enqueued close the channel

	userwg.Add(userCount)
	// take one user from the queue at a time
	for u := range users {
		// find them a computer
		c := <-computers
		// let them use the computer
		go u.start(c)
	}
	userwg.Wait()
	fmt.Println("All users have finished using the cafe. Shutting down!")
}
