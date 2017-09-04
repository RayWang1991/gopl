package main

import (
	"fmt"
	"bytes"
)

type treeNode struct {
	value       int
	left, right *treeNode
}

func (t *treeNode)String()string{
	if t == nil {
		return "nil"
	}
	buf := bytes.Buffer{}
	buf.WriteByte('(')
	buf.WriteString(fmt.Sprintf("%d ",t.value))
	buf.WriteString(t.left.String())
	buf.WriteByte(' ')
	buf.WriteString(t.right.String())
	buf.WriteByte(')')
	return buf.String()
}

func main() {
	values := []int{13, 1, 4, 21, 1, 0}
	fmt.Println(values)
	values = treesort(values)
	fmt.Println(values)

	root := &treeNode{0, &treeNode{value: 1}, &treeNode{value: 2}}
	//root := (*treeNode)(nil)
	fmt.Printf("LISP Printer:%v\n", root)
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
