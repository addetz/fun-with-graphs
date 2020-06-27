package main

import (
	"container/heap"

	utils "github.com/fun-with-graphs/puzzles"
)

/***
The question we'll work through is the following:
return a new sorted merged list from K sorted lists, each with size N.

For example, if we had [[10, 15, 30], [12, 15, 20], [17, 20, 32]],
the result should be [10, 12, 15, 15, 17, 20, 20, 30, 32].
Here, we can see that we only need to look at K elements in each of the lists
to find the smallest element initially. Heaps are great for finding the smallest
element. Let's say the smallest element is E.
Once we get E, we know we're interested in only the next element of
the list that held E. Then we'd extract out the second smallest element and etc.
The time complexity for this would be O(KN log K),
since we remove and append to the heap K * N times.
*/

func main() {
	input := [][]int{{10, 15, 30}, {12, 15, 20}, {17, 20, 32}}
	output := make([]int, 0)
	tHeap := &tupleHeap{}

	//Initialize the heap.
	heap.Init(tHeap)
	fcol := initFirstCol(input)
	for _, f := range fcol {
		heap.Push(tHeap, f)
	}

	// While the heap is not empty we need to:
	// Extract the minimum element from the heap: (value, list index, element index)
	// If the element index is not at the last index, add the next tuple in the list index.
	currentIndex := 0
	for tHeap.Len() != 0 {
		//Take min elem from the list
		p := heap.Pop(tHeap)
		cval := p.(*tuple)
		output = append(output, cval.val)
		//Is there another elem at this row?
		if cval.col < len(input[cval.row])-1 {
			nelem := &tuple{
				val: input[cval.row][cval.col+1],
				row: cval.row,
				col: cval.col + 1,
			}
			heap.Push(tHeap, nelem)
		}
		currentIndex++
	}

	utils.PrintIntMatrix("Input", input)
	utils.PrintInt("Sorted output", output)
}

func initFirstCol(input [][]int) []*tuple {
	fcol := make([]*tuple, len(input))
	for i, v := range input {
		fcol[i] = &tuple{
			val: v[0],
			row: i,
			col: 0,
		}
	}
	return fcol
}
