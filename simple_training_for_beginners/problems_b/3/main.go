// https://codeforces.com/problemset/problem/1399/B

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
		var n int
		fmt.Fscan(in, &n)

		sweets := make([]int64, n)
		oranges := make([]int64, n)
		minSweets := int64(1e9) + 1
		minOranges := int64(1e9) + 10
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &sweets[i])
			if sweets[i] < minSweets {
				minSweets = sweets[i]
			}
		}
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &oranges[i])
			if oranges[i] < minOranges {
				minOranges = oranges[i]
			}
		}

		var actions, a, b int64
		for i := 0; i < n; i++ {
			a = sweets[i] - minSweets
			b = oranges[i] - minOranges
			actions += max(a, b)
		}

		fmt.Println(actions)
	}
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
