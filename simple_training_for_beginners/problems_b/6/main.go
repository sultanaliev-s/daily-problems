// https://codeforces.com/problemset/problem/1360/B

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var t int
	fmt.Fscan(in, &t)

	for ; t > 0; t-- {
		var n int
		fmt.Fscan(in, &n)
		sportsmen := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &sportsmen[i])
		}

		sort.Ints(sportsmen)

		minDiff := 10000
		for i := 1; i < n; i++ {
			if minDiff > abs(sportsmen[i]-sportsmen[i-1]) {
				minDiff = abs(sportsmen[i] - sportsmen[i-1])
			}
		}

		fmt.Println(minDiff)
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
