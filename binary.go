package main

import (
	"fmt"
)

var count int

type Node struct {
	Key      int
	Value    int
	Left_Ch  *Node
	Right_Ch *Node
}

func main() {
	fmt.Println("binary tree")
	tree := &Node{Key: 23, Value: 56}

	tree.Insert(3)

	tree.Insert(34)

	fmt.Println(tree.Search(3))
	fmt.Println(count)

	fmt.Println(tree.Search(334))
	tree.Insert(23)
	fmt.Println(tree)

}

//search method
func (n *Node) Search(i int) bool {

	count++
	if n == nil {
		return false
	}
	if n.Key < i {
		// go right

		return n.Right_Ch.Search(i)

	} else if n.Key > i {
		//go left

		return n.Left_Ch.Search(i)

	}
	return true

}

//insert method
func (n *Node) Insert(i int) {
	if n == nil {
		n = &Node{Key: i, Value: 11}
	}
	if n.Key < i {
		// right insert
		if n.Right_Ch == nil {
			n.Right_Ch = &Node{Key: i, Value: 12}
		} else {
			n.Right_Ch.Insert(i)
		}
	} else if n.Key > i {
		if n.Left_Ch == nil {
			n.Left_Ch = &Node{Key: i, Value: 2}
		} else {
			n.Left_Ch.Insert(i)
		}
	} else {
		n.Value = 77777777777777
	}
}
