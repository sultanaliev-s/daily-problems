// https://codeforces.com/problemset/problem/1213/C

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
		var n, m int64
		fmt.Fscan(in, &n, &m)
		var res, tens, x int64
		x = m
		for x <= m*10 {
			tens += x % 10
			x += m
		}

		p := n / m / 10
		res += tens * p

		x = m
		for i := int64(0); i < n/m%10; i++ {
			res += x % 10
			x += m
		}

		fmt.Fprintln(out, res)
	}
}
