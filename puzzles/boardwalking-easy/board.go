package main

type board struct {
	points map[string]point
	v      *visiter
}

func newBoard() board {
	return board{
		points: make(map[string]point),
		v:      newVisiter(),
	}
}

func (b *board) addPoint(x, y int) {
	p, key := newPoint(x, y, false)
	b.points[key] = p
}

func (b *board) addWall(x, y int) {
	p, key := newPoint(x, y, true)
	b.points[key] = p
}

func (b *board) addDestination(x, y int) {
	p, key := newPoint(x, y, false)
	p.isDestination = true
	b.points[key] = p
}

func (b *board) addStart(x, y int) {
	b.v.push(generatePointKey(x, y), 0)
}

func (b board) hasNext() bool {
	return !b.v.isEmpty()
}

func (b board) next() vnode {
	return b.v.pop()
}

func (b board) nextSteps(n vnode) {
	p, ok := b.points[n.key]
	if !ok {
		return
	}
	cands := p.getNeighbors()
	for _, c := range cands {
		cp, ok := b.points[c]
		if !ok {
			continue
		}
		if cp.isWall {
			continue
		}
		b.v.push(c, n.distance+1)
	}
}

func (b board) isTarget(key string) bool {
	p, ok := b.points[key]
	if !ok {
		return false
	}
	return p.isDestination
}
