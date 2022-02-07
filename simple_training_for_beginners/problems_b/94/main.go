// https://codeforces.com/problemset/problem/1345/B

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
	for ; t > 0; t-- {
		var n int
		fmt.Fscan(in, &n)

		res := 0
		for n > 1 {
			h := 0
			t := 0
			for {
				c := count(h)
				if c+t > n {
					break
				}
				t += c
				h++
			}
			n -= t
			res++
		}

		fmt.Println(res)
	}
}

func count(h int) int {
	return h + ((h + 1) * 2)
}
