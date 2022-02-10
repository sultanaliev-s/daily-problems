// https://codeforces.com/problemset/problem/1353/C

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
		var n int64
		fmt.Fscan(in, &n)
		var res, c, l int64
		res, c, l = 0, 1, 1
		for c < n {
			res += (c*4 + 4) * l
			c += 2
			l++
		}

		fmt.Fprintln(out, res)
	}
}
