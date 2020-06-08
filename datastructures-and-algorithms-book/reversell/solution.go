package main

import "fmt"

type node struct {
    next *node
    val rune
}
func main() {
    head := &node{nil, 'a'}
    head.next = &node{nil, 'b'}
    head.next.next = &node{nil, 'c'}
    printList("Original", head)

    revHead := reverse(head)
    printList("Reversed", revHead)
}

func printList(title string, head *node) {
    fmt.Println(title)
    curr := head
    for {
        if curr == nil {
            break
        }
        fmt.Printf("%v \n", curr)
        curr = curr.next
    }
}

func reverse(head *node) *node {
    curr := head
    var prevNode *node
    for {
        if curr == nil {
            break
        }
        nextNode := curr.next
        // reverse by making next be prev
        curr.next = prevNode
        // move prev to curr
        prevNode = curr
        // move curr to next
        curr = nextNode
    }
    // the reversed head is prev as curr will be nil
    return prevNode
}
