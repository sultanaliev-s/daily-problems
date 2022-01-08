// https://codeforces.com/problemset/problem/900/A

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

	var neg, pos int
	for i, x, y := 0, 0, 0; i < n; i++ {
		fmt.Fscan(in, &x, &y)
		if x < 0 {
			neg++
		} else {
			pos++
		}
	}

	if neg <= 1 || pos <= 1 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
