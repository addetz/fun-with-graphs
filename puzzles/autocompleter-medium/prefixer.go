package main

import (
	"strings"
)

type node struct {
	val       string
	children  map[string]*node
	isWordEnd bool // sometimes a node may denote a full word but at the same time have children
}

func newNode() *node {
	return &node{
		children: make(map[string]*node),
	}
}

type prefixer struct {
	root *node
}

func newPrefixer() prefixer {
	return prefixer{root: newNode()}
}

func (p *prefixer) insert(word string) {
	words := strings.Split(word, "")
	current := p.root
	for _, w := range words {
		_, ok := current.children[w]
		if !ok {
			current.children[w] = newNode()
		}
		current = current.children[w]
	}
	current.val = word
	current.isWordEnd = true
}

func (p *prefixer) searchForPrefix(s string) []string {
	chars := strings.Split(s, "")
	current := p.root
	for _, c := range chars {
		next, ok := current.children[c]
		if !ok { // nothing found for the letter
			current = nil
			break
		}
		current = next
	}
	if current == nil {
		return []string{} // nothing found for the prefix
	}
	// traverse the children until you find the complete words
	results := make([]string, 0)
	nodes := make([]*node, 0)
	nodes = append(nodes, current)
	for {
		// exit if there are no more nodes to traverse
		if len(nodes) == 0 {
			break
		}
		// pop the head node
		head := nodes[0]
		nodes = nodes[1:]
		// append if the node is end
		if head.isWordEnd {
			results = append(results, head.val)
		}
		//enqueue children to the end
		for _, c := range head.children {
			if c != nil {
				nodes = append(nodes, c)
			}
		}
	}
	return results
}
