package main

import "fmt"

/**
Given two singly linked lists that intersect at some point,
find the intersecting node. The lists are non-cyclical.

For example, given A = 3 -> 7 -> 8 -> 10 and
B = 99 -> 1 -> 8 -> 10, return the node with value 8.

In this example, assume nodes with the same value are the exact same node objects.

Do this in O(M + N) time (where M and N are the lengths of the lists)
and constant space.
*/
type list struct {
	root  *node
	count int
}

func (l list) traverse(count int) *node {
	trav := l.root
	for i := 0; i < count; i++ {
		trav = trav.next
	}
	return trav
}

type node struct {
	value int
	next  *node
}

func main() {
	a := build(3, 7, 8, 10)
	b := build(100, 99, 1, 8, 10)
	aroot := a.root
	broot := b.root
	if a.count > b.count {
		aroot = a.traverse(a.count - b.count)
	}
	if b.count > a.count {
		broot = b.traverse(b.count - a.count)
	}
	intersection := findIntersection(aroot, broot)
	if intersection == nil {
		fmt.Println("The lists do not intersect")
		return
	}
	fmt.Printf("Intersection is %d \n", intersection.value)
}

func findIntersection(a *node, b *node) *node {
	if a == nil { //both  lists must be at the end since they were equal length
		return nil
	}
	if a.value == b.value {
		return a
	}
	return findIntersection(a.next, b.next)
}

func build(values ...int) *list {
	root := &node{
		value: values[0],
	}
	prev := root
	count := 1
	for i := 1; i < len(values); i++ {
		prev.next = &node{value: values[i]}
		prev = prev.next
		count++
	}
	return &list{
		root:  root,
		count: count,
	}
}
