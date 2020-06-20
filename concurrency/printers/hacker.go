package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	jobMaxSeconds  = 10
	maxWaitSeconds = 5 * time.Second
)

type hacker struct {
	name                 string
	remainingOperations  int
	successfulOperations int
	wg                   *sync.WaitGroup
	jobs                 chan<- job
	result               <-chan string
	done                 <-chan struct{}
}

func newHacker(name string, jobs chan<- job, results <-chan string,
	done <-chan struct{}, wg *sync.WaitGroup) *hacker {
	return &hacker{
		name:   name,
		done:   done,
		wg:     wg,
		jobs:   jobs,
		result: results,
	}
}

func (h *hacker) start() {
	fmt.Printf("Hacker %s has started ... \n", h.name)
	defer h.wg.Done()
	idleTimer := time.NewTimer(maxWaitSeconds)
	for {
		select {
		case <-h.done:
			fmt.Printf("Hackathon has ended and hacker %s has quit.\n", h.name)
			return
		case r := <-h.result:
			// job has completed
			fmt.Printf("Job %s has completed. \n", r)
			h.successfulOperations++
			// as soon as one job ends we go back to idle as we should be doing more work
			idleTimer.Reset(maxWaitSeconds)
		case h.jobs <- h.getPrintJob():
			fmt.Printf("Hacker %s successfully sent job.\n", h.name)
			// stop the idle timer on a successful send as the job may take longer than timeout
			idleTimer.Stop()
		case <-idleTimer.C:
			fmt.Printf("Hacker %s rage quit due to expired timer.\n", h.name)
			return
		}
	}
}

func (h *hacker) getPrintJob() job {
	return job{
		name: fmt.Sprintf("%s-job-%d", h.name, h.successfulOperations+1),
		d:    time.Duration(rand.Intn(jobMaxSeconds)) * time.Second,
		done: h.done,
	}
}
