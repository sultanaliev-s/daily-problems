// https://codeforces.com/problemset/problem/1389/C

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
		var s string
		fmt.Fscan(in, &s)

		var res int
		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				res = max(res, solve(s, i, j))
			}
		}

		fmt.Fprintln(out, len(s)-res)
	}
}

func solve(s string, x, y int) int {
	res := 0

	for i := range s {
		if int(s[i]-'0') == x {
			res++
			x, y = y, x
		}
	}

	if x != y && res%2 == 1 {
		res--
	}

	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
