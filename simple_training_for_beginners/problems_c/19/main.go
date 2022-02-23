// https://codeforces.com/problemset/problem/1221/C

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
		var c, m, lalki int
		fmt.Fscan(in, &c, &m, &lalki)

		x := min(c, m)
		lalki += c - x + m - x
		l := 0
		r := x
		for l < r {
			m := (l + r) / 2
			q := x - m
			if m < lalki+(q*2)-1 {
				l = m + 1
			} else {
				r = m
			}
		}

		g := x - l
		res := min(l, lalki+g*2)
		fmt.Fprintln(out, res)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
