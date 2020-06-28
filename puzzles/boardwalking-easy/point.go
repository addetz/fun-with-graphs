package main

import "fmt"

type point struct {
	x, y          int
	isWall        bool
	isDestination bool
}

func newPoint(x, y int, wall bool) (point, string) {
	return point{
		x:      x,
		y:      y,
		isWall: wall,
	}, generatePointKey(x, y)
}

func (p point) getNeighbors() []string {
	return []string{
		generatePointKey(p.x, p.y+1), //up
		generatePointKey(p.x, p.y-1), //down
		generatePointKey(p.x-1, p.y), //left
		generatePointKey(p.x+1, p.y), //right
	}
}

func generatePointKey(x, y int) string {
	return fmt.Sprintf("%d-%d", x, y)
}
