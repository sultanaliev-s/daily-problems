// https://codeforces.com/problemset/problem/1136/B

package main

import (
	"bufio"
	"fmt"
	"os"
)

type Edge struct {
	u int
	v int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	var n, k int
	fmt.Fscan(in, &n, &k)

	stoneThrows := n + 1
	shortSide := min(k-1, n-k)
	res := stoneThrows + shortSide + (n - 1) + n

	fmt.Fprintln(out, res)
	out.Flush()

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
