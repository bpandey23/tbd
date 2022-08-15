package main

import (
	"fmt"
)

type Node struct {
	Value int
	Next  *Node
}

func main() {
	fmt.Println("link list stack implementation")
	head := CreateNewNode(30)
	head.Next = CreateNewNode(49)
	head.Next.Next = CreateNewNode(50)
	TraverseLL(head)
}

func CreateNewNode(value int) *Node {
	var node Node
	node.Value = value
	node.Next = nil
	return &node
}

func TraverseLL(head *Node) {
	temp := head
	if temp == nil {
		return
	}

	for temp != nil {
		fmt.Println(temp.Value)
		temp = temp.Next
	}
}
