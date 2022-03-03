// https://codeforces.com/problemset/problem/1023/C

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var n, k int
	var s string
	fmt.Scan(&n, &k, &s)

	toDel := (n - k) / 2
	open := 0
	close := 0
	for i := range s {
		if s[i] == '(' && open < toDel {
			open++
		} else if s[i] == ')' && close < toDel {
			close++
		} else {
			fmt.Fprintf(out, "%c", s[i])
		}
	}
	fmt.Fprintln(out)
}
