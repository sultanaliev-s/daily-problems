// https://codeforces.com/problemset/problem/1257/B

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
		var x, y int
		fmt.Fscan(in, &x, &y)

		if (x > 3) || (x == 2 && y <= 3) || x >= y {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
