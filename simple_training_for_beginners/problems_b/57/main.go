// https://codeforces.com/problemset/problem/1054/B

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

	max := -1
	x := 0
	isPossible := true
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &x)
		if max+1 >= x {
			if max < x {
				max = x
			}
		} else {
			fmt.Println(i)
			isPossible = false
			break
		}
	}

	if isPossible {
		fmt.Println(-1)
	}
}
