package main

import (
	"fmt"
)

func main() {
	fmt.Println("union find")
	n := 100
	id := make([]int, n)
	//size list
	sz := make([]int, n)

	//initializing the list for quick find
	for i := 0; i < n; i++ {
		id[i] = i
		sz[i] = 1
	}
	//start
	fmt.Println("is 3 ,77 connected : ", Connected(3, 77, id))
	Union(3, 77, id, sz)
	//fmt.Println(id)
	fmt.Println("is 3 ,77 connected : ", Connected(3, 77, id))
	fmt.Println("is 4 ,8 connected : ", Connected(4, 8, id))
	Union(4, 8, id, sz)
	fmt.Println("is 4 ,8 connected : ", Connected(4, 8, id))
}

//finding root of the number
func RootOfNumber(i int, id []int) (root int) {
	for i != id[i] {
		//doing path compression of type single one pass variant
		id[i] = id[id[i]]
		i = id[i]
	}
	return i
}

//to check if two numbers are connected
func Connected(p, q int, id []int) bool {
	//check if two numbers are connected
	if RootOfNumber(p, id) == RootOfNumber(q, id) {
		return true
	} else {
		return false
	}
}

//to do union of two numbers
func Union(p, q int, id, sz []int) {
	i := id[p]
	j := id[q]
	if i == j {
		return
	}
	if sz[i] < sz[j] {
		id[i] = j
		sz[j] = sz[j] + sz[i]
	} else {
		id[j] = i
		sz[i] = sz[i] + sz[j]
	}

}
