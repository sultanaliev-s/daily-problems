// https://codeforces.com/problemset/problem/935/B

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n int
	var path string
	fmt.Fscan(in, &n, &path)

	var coins, x, y int
	for i := range path {
		if path[i] == 'U' {
			x++
			if x-y == 1 && i > 1 && path[i-1] == path[i] {
				coins++
			}
		} else {
			y++
			if y-x == 1 && i > 1 && path[i-1] == path[i] {
				coins++
			}
		}
	}

	fmt.Println(coins)
}
