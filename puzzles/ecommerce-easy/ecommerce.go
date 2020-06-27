package main

import "fmt"

/**
You run an e-commerce website and want to record the last N order ids in a log.
Implement a data structure to accomplish this, with the following API:
record(order_id): adds the order_id to the log
get_last(i): gets the ith last element from the log.
i is guaranteed to be smaller than or equal to N.
You should be as efficient with time and space as possible.
*/
const (
	n int = 5
)

type ordersLog struct {
	orders       [n]int
	currentCount int
}

func (o *ordersLog) record(id int) {
	o.orders[o.currentCount%n] = id
	o.currentCount++
}

func (o ordersLog) getLast(i int) int {
	if i > o.currentCount {
		return 0
	}
	// getLast(1) will give you the last item
	index := (o.currentCount - i) % n
	return o.orders[index]
}

func main() {
	olog := new(ordersLog)
	for i := 0; i < n; i++ {
		olog.record(i)
	}
	fmt.Printf("Last 1 without overflow: %d \n", olog.getLast(1))
	fmt.Printf("Last 2 without overflow: %d \n", olog.getLast(2))

	olog.record(n + 1)
	fmt.Printf("Last 1 with overflow: %d \n", olog.getLast(1))
	fmt.Printf("Last 2 with overflow: %d \n", olog.getLast(2))
}
