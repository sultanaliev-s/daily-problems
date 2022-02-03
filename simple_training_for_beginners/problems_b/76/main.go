// https://codeforces.com/problemset/problem/465/B

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(in, &n)

	var prev, cur, res int
	fmt.Fscan(in, &prev)
	if prev == 1 {
		res++
	}
	for i := 1; i < n; i++ {
		fmt.Fscan(in, &cur)
		if cur == 1 {
			if res > 0 && prev == 0 {
				res++
			}
			res++
		}
		prev = cur
	}

	fmt.Println(res)
}
