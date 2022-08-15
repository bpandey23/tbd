package main

import (
	"fmt"
)

func main() {
	fmt.Println("merge sort")

	arr := []int{2, 3, 4, 5, 4, 6, 56, 7, 34, 325, 47, 6, 34, 45, 57, 8, 34, 354, 65, 7, 57, 8, 83}

	sorted_arr := mergeSort(arr)
	fmt.Println(sorted_arr)
}

func mergeSort(arr []int) (sorted []int) {
	if len(arr) == 1 {
		return arr
	}
	first_half := mergeSort(arr[:len(arr)/2])
	second_half := mergeSort((arr[len(arr)/2:]))
	return merge(first_half, second_half)
}

func merge(farr, sarr []int) (sor_arr []int) {
	i, j := 0, 0
	for k := 0; k < len(farr)+len(sarr); k++ {
		if i == len(farr) {
			sor_arr = append(sor_arr, sarr[j])
			j++
		} else if j == len(sarr) {
			sor_arr = append(sor_arr, farr[i])
			i++
		} else if farr[i] > sarr[j] {
			sor_arr = append(sor_arr, sarr[j])
			j++
		} else {
			sor_arr = append(sor_arr, farr[i])
			i++
		}
	}
	return sor_arr
}
