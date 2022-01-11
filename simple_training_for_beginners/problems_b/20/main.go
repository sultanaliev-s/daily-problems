// https://codeforces.com/problemset/problem/1183/B

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
		var n, k int32
		fmt.Fscan(in, &n, &k)

		var min, max, x int32
		min = 1000 * 1000 * 1000
		for i := int32(0); i < n; i++ {
			fmt.Fscan(in, &x)
			if x > max {
				max = x
			}
			if x < min {
				min = x
			}
		}

		if max-min > k*2 {
			fmt.Println(-1)
		} else {
			fmt.Println(min + k)
		}
	}
}
