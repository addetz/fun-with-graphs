package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	minUsageTime = 5
	maxUsageTime = 20
)

type user struct {
	name string
	wg   *sync.WaitGroup
	out  chan<- computer // channel to return the computer back on when finished using
}

func newUser(name string, wg *sync.WaitGroup, out chan<- computer) user {
	return user{
		name: name,
		wg:   wg,
		out:  out,
	}
}

func (u user) start(c computer) {
	defer u.wg.Done()
	d := getUserDuration()
	fmt.Printf("%s has got %s for usage of %ds.\n", u.name, c.name, int(d.Seconds()))
	time.Sleep(d)
	fmt.Printf("%s finished using %s.\n", u.name, c.name)
	u.out <- c // user finished using the computer, return it
}

func getUserDuration() time.Duration {
	return time.Duration(rand.Intn(maxUsageTime-minUsageTime)+minUsageTime) * time.Second
}
