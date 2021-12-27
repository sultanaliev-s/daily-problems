// https://codeforces.com/problemset/problem/721/A

package main

import (
	"fmt"
)

func main() {
	var n int
	var line string
	fmt.Scanf("%d\n%s\n", &n, &line)

	groups := make([]int, n)

	b := line[0] == 'B'
	j := 0
	if b {
		groups[j]++
	}
	for i := 1; i < n; i++ {
		if line[i] != 'B' && b {
			b = false
			j++
		} else if line[i] == 'B' {
			b = true
			groups[j]++
		}
	}
	if line[len(line)-1] == 'B' {
		j++
	}

	fmt.Print(j)
	for i := 0; i < j; i++ {
		if i > 0 {
			fmt.Print(" ")
		} else {
			fmt.Println()
		}
		fmt.Print(groups[i])
	}
	fmt.Println()
}
