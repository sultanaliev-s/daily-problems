// https://codeforces.com/problemset/problem/716/A

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, c int
	fmt.Scanf("%d %d\n", &n, &c)
	in := bufio.NewReader(os.Stdin)

	words := 0
	prev := 0
	var cur int
	for i := 0; i < n; i++ {
		fmt.Fscanf(in, "%d", &cur)
		if cur-prev > c {
			words = 1
		} else {
			words++
		}

		prev = cur
	}

	fmt.Println(words)
}
