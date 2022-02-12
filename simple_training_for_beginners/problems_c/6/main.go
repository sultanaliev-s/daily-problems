// https://codeforces.com/problemset/problem/1335/C

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var t int
	fmt.Fscan(in, &t)
	for ; t > 0; t-- {
		var n int
		fmt.Fscan(in, &n)

		a := make([]int, n+1)
		var x int
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &x)
			a[x]++
		}

		diff, max := 0, 0
		for i := range a {
			if a[i] > 0 {
				diff++
			}
			if a[max] < a[i] {
				max = i
			}
		}

		maxCount := a[max]
		if diff == maxCount {
			fmt.Fprintln(out, diff-1)
		} else if diff < maxCount {
			fmt.Fprintln(out, diff)
		} else {
			fmt.Fprintln(out, maxCount)
		}
	}
}
