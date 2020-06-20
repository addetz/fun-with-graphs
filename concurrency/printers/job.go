package main

import (
	"fmt"
	"time"
)

type job struct {
	name string
	d    time.Duration
	done <-chan struct{}
}

func (j *job) print(pname string) {
	fmt.Printf("Printing job %s on printer %s started and will take %ds.\n", j.name, pname, j.d/time.Second)
	after := time.After(j.d)
	for {
		select {
		case <-j.done:
			fmt.Printf("Job %s was cancelled and exited.\n", j.name)
			return
		case <-after:
			fmt.Printf("Printing job %s on printer %s completed.\n", j.name, pname)
			return
		}
	}
}
