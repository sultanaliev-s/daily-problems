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
		var n, k, m int64
		fmt.Fscan(in, &n, &k)
		m = n - 1

		if n == 4 && k == 3 {
			fmt.Fprintln(out, -1)
		} else if m == k {
			p := m ^ 3
			fmt.Fprintln(out, m-1, m)
			fmt.Fprintln(out, 1, 3)
			fmt.Fprintln(out, 0, p)
			for i := int64(2); i < n/2; i++ {
				if i == p || i == 3 {
					continue
				}
				fmt.Fprintln(out, i, m-i)
			}
		} else {
			p := m ^ k
			fmt.Fprintln(out, k, m)
			if k != 0 {
				fmt.Fprintln(out, 0, p)
			}

			for i := int64(1); i < n/2; i++ {
				if i == p || i == k {
					continue
				}
				fmt.Fprintln(out, i, m-i)
			}
		}
	}
}
