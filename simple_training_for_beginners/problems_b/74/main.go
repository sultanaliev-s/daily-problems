// https://codeforces.com/problemset/problem/616/B

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n, m int
	fmt.Fscan(in, &n, &m)

	gr := make([][]int, n)
	for i := range gr {
		gr[i] = make([]int, m)
		for j := range gr[i] {
			fmt.Fscan(in, &gr[i][j])
		}
	}

	max := 0
	for i := 0; i < n; i++ {
		min := 1_000_000_001
		for j := 0; j < m; j++ {
			if gr[i][j] < min {
				min = gr[i][j]
			}
		}
		if min > max {
			max = min
		}
	}

	fmt.Println(max)
}
