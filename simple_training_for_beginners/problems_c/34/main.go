// https://codeforces.com/problemset/problem/519/C

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
	var n, m int
	fmt.Fscan(in, &n, &m)

	teams := 0
	for n > 0 && m > 0 && n+m > 2 {
		if n < m {
			n -= 1
			m -= 2
		} else {
			n -= 2
			m -= 1
		}
		teams++
	}

	fmt.Fprintln(out, teams)
}
