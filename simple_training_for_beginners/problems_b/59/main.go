// https://codeforces.com/problemset/problem/1008/B

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

	var x, y int
	fmt.Fscan(in, &x, &y)
	prev := max(x, y)
	isPossible := true
	for i := 1; i < n; i++ {
		fmt.Fscan(in, &x, &y)
		if prev >= x && prev >= y {
			prev = max(x, y)
		} else if prev >= x {
			prev = x
		} else if prev >= y {
			prev = y
		} else {
			isPossible = false
		}
	}

	if isPossible {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
