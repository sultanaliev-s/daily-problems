// https://codeforces.com/problemset/problem/702/A

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

	maxLen := 1
	l := 1
	var prev, cur int
	fmt.Fscan(in, &prev)
	for i := 1; i < n; i++ {
		fmt.Fscan(in, &cur)
		if cur > prev {
			l++
		} else {
			if maxLen < l {
				maxLen = l
			}
			l = 1
		}

		prev = cur
	}
	if maxLen < l {
		maxLen = l
	}

	fmt.Println(maxLen)
}
