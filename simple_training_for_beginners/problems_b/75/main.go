// https://codeforces.com/problemset/problem/499/B

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

	dict := make(map[string]string)
	var a, b string
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &a, &b)
		dict[a] = b
	}

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a)
		if i > 0 {
			fmt.Fprint(out, " ")
		}
		if len(a) <= len(dict[a]) {
			fmt.Fprint(out, a)
		} else {
			fmt.Fprint(out, dict[a])
		}
	}
}
