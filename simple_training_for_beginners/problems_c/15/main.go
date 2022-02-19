// https://codeforces.com/problemset/problem/1328/C

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
		var x string
		fmt.Fscan(in, &n, &x)

		a := make([]byte, n)
		b := make([]byte, n)

		aIsBigger := false
		for i := range x {
			if !aIsBigger {
				if x[i] == '1' {
					a[i] = x[i]
					b[i] = '0'
					aIsBigger = true
				} else if x[i] == '2' {
					a[i] = '1'
					b[i] = '1'
				} else {
					a[i] = '0'
					b[i] = '0'
				}
			} else {
				a[i] = '0'
				b[i] = x[i]
			}
		}

		fmt.Fprintln(out, string(a))
		fmt.Fprintln(out, string(b))
	}
}
