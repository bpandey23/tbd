package main

import (
	"fmt"
)

func main() {
	fmt.Println("brute-force_substring-search")
	txt := "brute-force_substring-search"
	pat := "string"
	fmt.Println(bf_string_search(pat, txt))

}

//this func returns the start and end of the substring(pat) if it exists in the input(txt)
//otherwise it returns -1,-1
func bf_string_search(pat, txt string) (start, end int, err error) {
	p_r := []rune(pat)
	t_r := []rune(txt)
	m := len(pat)
	n := len(txt)
	fmt.Println(p_r, t_r, m, n)
	for i := 0; i < n-m; i++ {
		var j int
		for j = 0; j < m; j++ {
			if t_r[i+j] != p_r[j] {
				break
			}

		}
		if j == m {
			return i, i + j, nil
		}
	}
	return -1, -1, fmt.Errorf("no match found")
}
