package main

import (
	"fmt"
)

func main() {
	fmt.Println("brute-force_with_backup_string_search")
	txt := "brute-force_with_backup_string_search_foo barr"
	pat := "bar"
	fmt.Println(brute_force_with_backup_string_search(pat, txt))
}

func brute_force_with_backup_string_search(pat, txt string) (start, end int, err error) {
	n := len(txt)
	m := len(pat)

	//commented rune conversion , due to its limited scope
	p_r := []byte(pat) // []rune(pat)
	t_r := []byte(txt) //[]rune(txt)

	var i, j int
	for i, j = 0, 0; i < n && j < m; i++ {
		if t_r[i] == p_r[j] {
			j++
		} else {
			i = i - j
			j = 0
		}
	}
	if j == m {
		return i - j, i, nil
	}
	return -1, -1, fmt.Errorf("no match found")
}
