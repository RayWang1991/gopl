package main

import "fmt"

type treeNode struct {
	value       int
	left, right *treeNode
}

func main() {
	values := []int{13, 1, 4, 21, 1, 0}
	fmt.Println(values)
	values = treesort(values)
	fmt.Println(values)
}

func treesort(values []int) []int {
	var root *treeNode
	for _, v := range values {
		root = addV(root, v)
	}
	values = inorderAdd(root, values[:0])
	return values
}

func inorderAdd(node *treeNode, values []int) []int {
	if node == nil {
		return values
	}
	values = inorderAdd(node.left, values)
	values = append(values, node.value)
	values = inorderAdd(node.right, values)
	return values
}

func addV(node *treeNode, v int) *treeNode {
	if nil == node {
		node = &treeNode{value: v}
	} else {
		if v < node.value {
			node.left = addV(node.left, v)
		} else {
			node.right = addV(node.right, v)
		}
	}
	return node
}
