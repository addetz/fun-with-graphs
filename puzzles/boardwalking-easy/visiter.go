package main

type vnode struct {
	key      string
	distance int
}
type visiter struct {
	toVisit []vnode
	visited map[string]struct{}
}

func newVisiter() *visiter {
	return &visiter{
		toVisit: make([]vnode, 0),
		visited: make(map[string]struct{}, 0),
	}
}

// enqueue to the back of visit queue - for BFS stack
func (v *visiter) push(s string, distance int) {
	_, ok := v.visited[s]
	if ok {
		return
	}
	v.toVisit = append(v.toVisit, []vnode{{
		key:      s,
		distance: distance,
	}}...)
}

// pop last move from the top - return bool to see if there was something to pop
func (v *visiter) pop() vnode {
	head := v.toVisit[0]
	v.toVisit = v.toVisit[1:]
	v.visited[head.key] = struct{}{}
	return head
}

func (v visiter) isEmpty() bool {
	return len(v.toVisit) == 0
}
