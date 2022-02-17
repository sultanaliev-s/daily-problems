// https://codeforces.com/problemset/problem/1352/C

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
		var n, k int
		fmt.Fscan(in, &n, &k)

		mults := (k - 1) / (n - 1)
		cur := n * mults
		x := k - (cur - mults)
		res := cur + x

		fmt.Fprintln(out, res)
	}
}
