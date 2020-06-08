package main

type tuple struct {
    val, row, col int
}
// integerHeap a type
type tupleHeap []*tuple

// tupleHeap method - gets the length of integerHeap
func (theap tupleHeap) Len() int { return len(theap) }

// tupleHeap method - checks if element of i index is less than j index
func (theap tupleHeap) Less(i, j int) bool { return theap[i].val < theap[j].val }
// tupleHeap method -swaps the element of i to j index
func (theap tupleHeap) Swap(i, j int) { theap[i], theap[j] = theap[j], theap[i] }

//tupleHeap method -pushes the item
func (theap *tupleHeap) Push(heapintf interface{}) {
    *theap = append(*theap, heapintf.(*tuple))
}
//tupleHeap method -pops the item from the heap
func (theap *tupleHeap) Pop() interface{} {
    var previous tupleHeap = *theap
    n := len(previous)
    x1 := previous[n-1]
    *theap = previous[0 : n-1]
    return x1
}