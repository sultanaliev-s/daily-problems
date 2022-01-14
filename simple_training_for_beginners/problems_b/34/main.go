// https://codeforces.com/problemset/problem/34/B

package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	tellies := make([]int, n)
	for i := range tellies {
		fmt.Scan(&tellies[i])
	}

	sort.Ints(tellies)

	profit := 0
	for i := 0; i < n && tellies[i] < 0 && m > 0; i++ {
		profit += -tellies[i]
		m--
	}

	fmt.Println(profit)
}
