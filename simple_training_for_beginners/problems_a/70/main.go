// https://codeforces.com/problemset/problem/595/A

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

	notSleeping := 0
	var x, y int
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Fscan(in, &x, &y)
			if x+y > 0 {
				notSleeping++
			}
		}
	}

	fmt.Println(notSleeping)
}
