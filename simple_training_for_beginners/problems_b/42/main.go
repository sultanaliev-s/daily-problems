// https://codeforces.com/problemset/problem/1266/B

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
		var n int64
		fmt.Fscan(in, &n)

		rem := n % 14
		if n > 14 && 1 <= rem && rem <= 6 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
