package main

import "fmt"

type Stack struct {
	items []int
}

//push
func (s *Stack) push(i int) {
	s.items = append(s.items, i)
}

func (s *Stack) pop() (i int) {
	if len(s.items) == 0 {
		fmt.Println("underflow")
		return -1
	}
	i = s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return
}

func main() {
	var myStack Stack
	fmt.Println(myStack)
	myStack.push(5)
	myStack.push(51)
	myStack.push(52)
	fmt.Println(myStack)
	fmt.Println(myStack.pop())
	fmt.Println(myStack)
	fmt.Println(myStack.pop())
	fmt.Println(myStack)
	fmt.Println(myStack.pop())
	fmt.Println(myStack)
	fmt.Println(myStack.pop())
	fmt.Println(myStack)
	fmt.Println(myStack.pop())
	fmt.Println(myStack)
}
