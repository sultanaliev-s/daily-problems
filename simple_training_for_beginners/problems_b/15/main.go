// https://codeforces.com/problemset/problem/1296/B

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
		var x int
		fmt.Fscan(in, &x)

		k := 1000000000
		res := 0
		for x > 0 {
			for k > x {
				k /= 10
			}
			res += k
			x -= k
			x += k / 10
		}

		fmt.Println(res)
	}
}
