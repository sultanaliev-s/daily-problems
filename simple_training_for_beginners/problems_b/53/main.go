// https://codeforces.com/problemset/problem/1107/B

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

	for i := 0; i < n; i++ {
		var k, x int64
		fmt.Fscan(in, &k, &x)

		res := (k-1)*9 + x
		fmt.Println(res)
	}
}
