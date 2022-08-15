package main

import (
	"fmt"
)

type Node struct {
	key        int
	value      int
	leftchild  *Node
	rightchild *Node
	colour     bool //colour of parent link
}

func (n *Node) RedlinkCheck() bool {
	if n == nil {
		n.colour = false
		return false
	}
	n.colour = true
	return true
}


//tree is right leaning at this moment
func (n *Node) LeftRotate(h *Node) {
	x := h.rightchild

	h.rightchild = x.leftchild
	x.leftchild = h
	x.colour = h.colour
	h.colour = true
	return

}


//tree is left leaning at this moment
func (n *Node)RightRotate(h *Node){
	x:=h.leftchild
	h.leftchild=x.rightchild
	
}

func (n *Node) Search(k int) (result bool, val int) {
	if n == nil {
		return false, -1
	}
	if n.key < k {
		//go right

		return n.rightchild.Search(k)

	} else if n.key > k {
		//go left

		return n.leftchild.Search(k)

	}

	return true, n.value
}

func main() {
	fmt.Println("left leaning red black tree")
	var tree Node
	//tree := &Node{key: 7, value: 78}
	//fmt.Println(tree)
	fmt.Println(tree.Search(7))
	fmt.Println(tree.Search(74))
}
