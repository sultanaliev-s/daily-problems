// https://codeforces.com/problemset/problem/822/B

package main

import (
	"fmt"
)

func main() {
	var n, m int
	var s, t string
	fmt.Scan(&n, &m, &s, &t)

	minN := 10000
	minPostions := make([]int, 0)

	for i := 0; i <= m-n; i++ {
		cnt := 0
		positions := make([]int, n)
		for j := 0; j < n; j++ {
			if s[j] != t[j+i] {
				positions[cnt] = j + 1
				cnt++
			}
		}

		if minN > cnt {
			minN = cnt
			minPostions = positions
		}
	}

	fmt.Println(minN)
	for i := 0; i < minN; i++ {
		if i != 0 {
			fmt.Print(" ")
		}
		fmt.Print(minPostions[i])
	}
	fmt.Println()
}
