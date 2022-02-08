// https://codeforces.com/problemset/problem/1213/B

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
	for ; t > 0; t-- {
		var n int
		fmt.Fscan(in, &n)

		a := make([]int, n)
		for i := range a {
			fmt.Fscan(in, &a[i])
		}

		min := a[n-1]
		bads := 0
		for i := n - 2; i >= 0; i-- {
			if a[i] > min {
				bads++
			} else {
				min = a[i]
			}
		}

		fmt.Println(bads)

	}
}
