// https://codeforces.com/problemset/problem/1325/B

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
		var n, x int
		fmt.Fscan(in, &n)
		nums := make(map[int]struct{})
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &x)
			nums[x] = struct{}{}
		}

		fmt.Println(len(nums))
	}
}
