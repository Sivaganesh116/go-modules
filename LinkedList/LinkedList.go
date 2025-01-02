package LinkedList

import (
	"fmt"
)

type Node struct {
	Next *Node
	Val int
}

func (curNode *Node) Insert (node *Node) {
	curNode.Next = node
}

func (root *Node) Print () {
	for root != nil {
		fmt.Printf("%v ", root.Val)
		root = root.Next
	}
}
