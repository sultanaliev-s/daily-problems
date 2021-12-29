// https://codeforces.com/problemset/problem/677/A

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n, h int
	fmt.Fscan(in, &n, &h)

	width := 0
	for i, x := 0, 0; i < n; i++ {
		fmt.Fscan(in, &x)
		if x > h {
			width++
		}
		width++
	}

	fmt.Println(width)
}
