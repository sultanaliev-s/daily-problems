// https://codeforces.com/problemset/problem/1358/B

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

		babushkas := make([]int, n)
		for i := range babushkas {
			fmt.Fscan(in, &babushkas[i])
		}

		sort.Slice(babushkas, func(i, j int) bool {
			return babushkas[i] > babushkas[j]
		})

		i := 0
		for ; i < n; i++ {
			if babushkas[i] <= n-i {
				break
			}
		}

		res := n - i + 1
		if i == n {
			res = 1
		}

		fmt.Println(res)
	}
}
