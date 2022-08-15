package main

import (
	"fmt"
)

type Node struct {
	Val  int
	Next *Node
}

type LinkL struct {
	Head *Node
	Len  int
}

func (l *LinkL) prepend(n *Node) {
	second := l.Head
	l.Head = n
	l.Head.Next = second
	l.Len += 1
}

func main() {
	fmt.Println("linked list")
	var ll LinkL
	no := &Node{Val: 2}
	fmt.Println(no)
	//fmt.Println(ll.Head.Val, ll.Len)
	ll.prepend(no)
	fmt.Println(ll.Head.Val, ll.Len)
	no = &Node{Val: 22}
	ll.prepend(no)
	fmt.Println(ll.Head.Val, ll.Len)
	no = &Node{Val: 2344}
	ll.prepend(no)
	fmt.Println(ll.Head.Val, ll.Len)
	no = &Node{Val: 45}
	ll.prepend(no)
	fmt.Println(ll.Head.Val, ll.Len)

}
