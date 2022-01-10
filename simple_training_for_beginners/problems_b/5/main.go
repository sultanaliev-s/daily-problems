// https://codeforces.com/problemset/problem/1391/B

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var t int
	fmt.Fscan(in, &t)

	var n, m int
	var line string
	for ; t > 0; t-- {
		fmt.Fscan(in, &n, &m)
		changes := 0
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &line)

			if i != n-1 {
				if line[m-1] == 'R' {
					changes++
				}
			} else {
				for j := 0; j < m; j++ {
					if line[j] == 'D' {
						changes++
					}
				}
			}

		}

		fmt.Println(changes)
	}
}
