// https://codeforces.com/problemset/problem/1311/C

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
		var s string
		fmt.Fscan(in, &n, &m, &s)

		m++
		miss := make([]int, n)
		alph := make([]int64, 26)
		var x int
		for i := 0; i < m-1; i++ {
			fmt.Fscan(in, &x)
			miss[x-1]++
		}
		miss[n-1]++

		h := int64(m)
		for i := range miss {
			alph[s[i]-'a'] += h
			h -= int64(miss[i])
		}

		for i := 0; i < 26; i++ {
			if i != 0 {
				fmt.Fprint(out, " ")
			}
			fmt.Fprint(out, alph[i])
		}
		fmt.Fprintln(out)
	}
}
