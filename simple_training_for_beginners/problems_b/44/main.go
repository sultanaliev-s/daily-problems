// https://codeforces.com/problemset/problem/1244/B

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var t int
	fmt.Fscan(in, &t)

	var n int
	var line string
	for ; t > 0; t-- {
		fmt.Fscan(in, &n, &line)

		l, r := 0, 0
		for i := 0; i < n; i++ {
			if line[i] == '1' {
				break
			}
			l++
		}
		for i := n - 1; i >= 0; i-- {
			if line[i] == '1' {
				break
			}
			r++
		}

		if r < l {
			l = r
		}

		var res int

		if l == n {
			res = n
		} else {
			res = 2*n - 2*l
		}

		fmt.Println(res)
	}
}
