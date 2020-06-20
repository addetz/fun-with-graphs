package main

import (
	"fmt"
	"sync"
)

type printer struct {
	name   string
	lock   sync.RWMutex
	jobs   <-chan job
	result chan<- string
	done   <-chan struct{}
	wg     *sync.WaitGroup
}

func newPrinter(name string, jobs <-chan job, result chan<- string, done <-chan struct{}, wg *sync.WaitGroup) *printer {
	return &printer{
		name:   name,
		lock:   sync.RWMutex{},
		jobs:   jobs,
		result: result,
		done:   done,
		wg:     wg,
	}
}

func (p *printer) start() {
	fmt.Printf("Printer %s has started ...\n", p.name)
	defer p.wg.Done()
	for {
		select {
		case <-p.done:
			fmt.Printf("Hackathon ended, closing printer %s \n", p.name)
			return
		case job := <-p.jobs:
			p.lock.Lock()
			job.print(p.name)
			// signal that the job has finished to whoever is listening
			p.result <- p.name
			p.lock.Unlock()
		}
	}
}
