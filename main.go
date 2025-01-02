package main

import (
	"fmt"
	"github.com/Sivaganesh116/go-modules/LinkedList"
)

func main() {
	var node *LinkedList.Node = &LinkedList.Node{nil, 10}

	fmt.Println(node.Val)
}