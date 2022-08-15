package main

import (
	"fmt"
)

type Que struct {
	Items []int
}

func (q *Que) enq(i int) {
	q.Items = append(q.Items, i)
}

func (q *Que) deq() (i int) {
	if len(q.Items) == 0 {
		fmt.Println("no element in q")
		return -1
	}
	i = q.Items[0]
	q.Items = q.Items[1:]
	return
}

func main() {
	fmt.Println("que")
	var qq Que
	fmt.Println(qq)
	qq.enq(455)
	qq.enq(4525)
	qq.enq(45)
	qq.enq(4)
	fmt.Println(qq)
	fmt.Println(qq.deq())

	fmt.Println(qq.deq())

	fmt.Println(qq.deq())

	fmt.Println(qq.deq())

	fmt.Println(qq.deq())
}
