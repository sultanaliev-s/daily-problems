// https://codeforces.com/problemset/problem/1343/C

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
		fmt.Fscan(in, &n)

		ar := make([]int, n)
		for i := range ar {
			fmt.Fscan(in, &ar[i])
		}

		var res int64
		var isPos bool
		for i := 0; i < n; i++ {
			isPos = IsPositive(ar[i])
			l := ar[i]
			for ; i < n; i++ {
				if IsPositive(ar[i]) != isPos {
					i -= 1
					break
				} else if ar[i] > l {
					l = ar[i]
				}
			}
			res += int64(l)
		}

		fmt.Fprintln(out, res)
	}
}

func IsPositive(x int) bool {
	return x > 0
}
