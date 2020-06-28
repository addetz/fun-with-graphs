package main

import (
	"fmt"
	"strings"
)

/**
Given the root to a binary tree, implement serialize(root),
which serializes the tree into a string,
and deserialize(s), which deserializes the string back into the tree.
*/
type node struct {
	val         string
	left, right *node
}

func main() {
	root := setupTree()
	ser := serialize(root)
	fmt.Printf("Serialized tree: %s \n", ser)
	deser := deserialize(ser)
	fmt.Printf("Deserialized tree root: %v \n", deser)
	fmt.Printf("Reserialized tree : %v \n", serialize(deser))
}

func deserialize(ser string) *node {
	values := strings.Split(strings.Trim(ser, " "), " ")
	var root *node
	// special handling for the root
	if values[0] != "nil" {
		root = &node{values[0], nil, nil}
	}
	current := root
	var direction bool // a flag to tell us which way to run
	nodes := make([]*node, 0)
	for i := 1; i < len(values); i++ {
		val := values[i]
		node := getNode(val)
		if !direction { // we should run left
			current.left = node
			nodes = append(nodes, node)
			direction = true
			continue
		}
		// we should run right
		current.right = node
		if node != nil { // append the node only if it's not nil
			nodes = append(nodes, node)
		}
		direction = false
		current = nodes[0] // completed run , move current node
		nodes = nodes[1:]
	}
	return root
}

func getNode(val string) *node {
	if val == "nil" {
		return nil
	}
	return &node{val, nil, nil}
}

func serialize(root *node) string {
	var sb strings.Builder
	// we will use in order traverse to serialize
	nodes := make([]*node, 0)
	nodes = append(nodes, root)
	for {
		// exit condition
		if len(nodes) == 0 {
			break
		}
		// pop the head
		current := nodes[0]
		nodes = nodes[1:]
		if current != nil {
			//write the node to the log
			sb.WriteString(fmt.Sprintf("%s ", current.val))
			nodes = append(nodes, current.left, current.right)
			continue
		}
		sb.WriteString("nil ")
	}

	return sb.String()
}

func setupTree() *node {
	ll := &node{"left.left", nil, nil}
	lr := &node{"left.right", nil, nil}
	l := &node{"left", ll, lr}
	rl := &node{"right.left", nil, nil}
	r := &node{"right", rl, nil}
	root := &node{"root", l, r}
	return root
}
