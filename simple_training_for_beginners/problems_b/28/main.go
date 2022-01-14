// https://codeforces.com/problemset/problem/1005/B

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var s, t string
	fmt.Fscan(in, &s, &t)

	similar := 0
	for i, j := len(s)-1, len(t)-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if s[i] != t[j] {
			break
		}
		similar++
	}

	steps := len(s) - similar + len(t) - similar
	fmt.Println(steps)
}
