// https://codeforces.com/problemset/problem/1326/B

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
	var n int
	fmt.Fscan(in, &n)

	var x, max int
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &x)
		if i != 0 {
			fmt.Fprint(out, " ")
		}
		fmt.Fprint(out, max+x)
		max = Max(max+x, max)
	}
	fmt.Fprintln(out)
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
