// https://codeforces.com/problemset/problem/1176/A

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var q int
	fmt.Fscan(in, &q)

	var n, a, b, c int64
	var count int
	for t := 0; t < q; t++ {
		fmt.Fscan(in, &n)
		a, b, c = n, n, n
		for n != 1 {
			if n%2 == 0 {
				a = n / 2
			}
			if n%3 == 0 {
				b = n / 3 * 2
			}
			if n%5 == 0 {
				c = n / 5 * 4
			}
			if a+b+c == n*3 {
				count = -1
				break
			}
			n = min(a, min(b, c))
			a, b, c = n, n, n
			count++
		}

		fmt.Println(count)
		count = 0
	}
}

func min(a, b int64) int64 {
	if a > b {
		return b
	}
	return a
}
