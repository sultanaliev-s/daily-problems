// https://codeforces.com/problemset/problem/1374/C

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
		var n int
		var s string
		fmt.Fscan(in, &n, &s)

		ops, cls, res := 0, 0, 0
		for i := range s {
			if s[i] == '(' {
				ops++
			} else {
				cls++
				if cls > ops {
					res++
					cls--
				}
			}
		}

		fmt.Fprintln(out, res)
	}
}
