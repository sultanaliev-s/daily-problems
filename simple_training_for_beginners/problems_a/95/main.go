// https://codeforces.com/problemset/problem/1385/B

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

	var n, x int
	var nums [51]bool
	for test := 0; test < t; test++ {
		fmt.Fscan(in, &n)
		n *= 2
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &x)
			if !nums[x] {
				if i != 0 {
					fmt.Print(" ")
				}
				fmt.Print(x)
				nums[x] = true
			} else {
				nums[x] = false
			}
		}
		fmt.Println()
	}
}
