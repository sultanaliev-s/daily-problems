// https://codeforces.com/problemset/problem/1366/C

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

		ar := make([][]int, n)
		for i := range ar {
			ar[i] = make([]int, m)
			for j := range ar[i] {
				fmt.Fscan(in, &ar[i][j])
			}
		}

		diags := n + m - 1
		type pair struct {
			zero int
			one  int
		}
		pairs := make([]pair, diags)
		for i := range ar {
			for j := range ar[i] {
				if ar[i][j] == 0 {
					pairs[i+j].zero++
				} else {
					pairs[i+j].one++
				}
			}
		}

		res := 0
		for i, j := 0, diags-1; i < j; i, j = i+1, j-1 {
			res += min(pairs[i].zero+pairs[j].zero, pairs[i].one+pairs[j].one)
		}

		fmt.Fprintln(out, res)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
