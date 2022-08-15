package main

import (
	"fmt"
)

func main() {
	fmt.Println("merge sort")
	arr := []int{3, 2, 3356, 45, 57, 87, 5, 356, -1, 47}

	//will call merge sort and it will return sorted array
	fmt.Println(mergeSort(arr))

}

func mergeSort(arr []int) (sort_array []int) {
	if len(arr) == 1 {
		return arr
	}
	first_half := mergeSort(arr[:len(arr)/2])
	second_half := mergeSort(arr[len(arr)/2:])
	return merge(first_half, second_half)

}

func merge(farr, sarr []int) []int {
	merged_arr := make([]int, 0)
	i, j := 0, 0
	for k := 0; k < (len(farr) + len(sarr)); k++ {
		if i == len(farr) {
			merged_arr = append(merged_arr, sarr[j])
			j++
		} else if j == len(sarr) {
			merged_arr = append(merged_arr, farr[i])
			i++
		} else if farr[i] > sarr[j] {
			merged_arr = append(merged_arr, sarr[j])
			j++
		} else {
			merged_arr = append(merged_arr, farr[i])
			i++
		}
	}
	return merged_arr
}
