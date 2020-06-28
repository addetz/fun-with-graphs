package main

import "fmt"

/**
You are given an M by N matrix consisting of booleans that represents a board.
Each True boolean represents a wall.
Each False boolean represents a tile you can walk on.

Given this matrix, a start coordinate, and an end coordinate,
return the minimum number of steps required to reach the end
coordinate from the start.
If there is no possible path, then return null.
You can move up, left, down, and right.
You cannot move through walls.
You cannot wrap around the edges of the board.
*/

func main() {
	start := []int{3, 0}
	end := []int{0, 0}
	board := setupBoard(start, end)
	steps, routePossible := getMinSteps(board)
	if !routePossible {
		fmt.Println("No route was possible.")
		return
	}
	fmt.Printf("Minimum number of steps: %d \n", steps)
}

func getMinSteps(b board) (int, bool) {
	// continue until the stack is empty
	for {
		// the queue is empty, break the loop
		if !b.hasNext() {
			break
		}
		// pop the top node
		pnode := b.next()
		// if it's the target we return its distance
		// since we are using BFS it will be found first
		if b.isTarget(pnode.key) {
			return pnode.distance, true
		}
		// otherwise generate the legal moves and append them
		b.nextSteps(pnode)
	}
	return 0, false
}

func setupBoard(start []int, end []int) board {
	b := newBoard()
	b.addPoint(0, 0)
	b.addPoint(0, 1)
	b.addPoint(0, 2)
	b.addPoint(0, 3)
	b.addWall(1, 0)
	b.addWall(1, 1)
	b.addPoint(1, 2)
	b.addWall(1, 3)
	b.addPoint(2, 0)
	b.addPoint(2, 1)
	b.addPoint(2, 2)
	b.addPoint(2, 3)
	b.addPoint(3, 0)
	b.addPoint(3, 1)
	b.addPoint(3, 2)
	b.addPoint(3, 3)
	b.addDestination(end[0], end[1])
	b.addStart(start[0], start[1])
	return b
}
