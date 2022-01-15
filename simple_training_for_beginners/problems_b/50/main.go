// https://codeforces.com/problemset/problem/1139/B

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

	chocs := make([]int64, n)
	for i := range chocs {
		fmt.Fscan(in, &chocs[i])
	}

	sum, prev := chocs[n-1], chocs[n-1]
	for i := n - 2; i >= 0; i-- {
		if chocs[i] < prev {
			sum += chocs[i]
			prev = chocs[i]
		} else if prev > 0 {
			sum += prev - 1
			prev = prev - 1
		}
	}

	fmt.Println(sum)
}
