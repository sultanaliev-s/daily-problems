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
		var n, m int
		fmt.Fscan(in, &n, &m)

		var l, r, x, res, mod, pt int64
		mod = 1_000_000_007
		pt = 1
		for i := 0; i < n-1; i++ {
			pt = pt * 2 % mod
		}
		for i := 0; i < m; i++ {
			fmt.Fscan(in, &l, &r, &x)
			res |= x
		}
		res = (res % mod * pt) % mod

		fmt.Fprintln(out, res)
	}
}
