package LinkedList

type Node struct {
	next *Node
	val int
}

func (curNode *Node) insert (node *Node) {
	curNode.next = node
}

func (root *Node) print () {
	for root != nil {
		fmt.Printf("%v ", root.val)
		root = root.next
	}
}