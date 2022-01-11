// https://codeforces.com/problemset/problem/1351/B

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
		var a1, b1, a2, b2 int
		fmt.Fscan(in, &a1, &b1, &a2, &b2)
		if a1 > b1 {
			a1, b1 = b1, a1
		}
		if a2 > b2 {
			a2, b2 = b2, a2
		}
		if a1+a2 == b1 && b1 == b2 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
