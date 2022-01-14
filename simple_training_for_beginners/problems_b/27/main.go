// https://codeforces.com/problemset/problem/1047/B

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

	var x, y, max int
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &x, &y)
		if max < x+y {
			max = x + y
		}
	}

	fmt.Println(max)
}
