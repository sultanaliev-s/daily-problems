// https://codeforces.com/problemset/problem/1300/B

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

		pupils := make([]int, n*2)
		for i := 0; i < n+n; i++ {
			fmt.Fscan(in, &pupils[i])
		}

		sort.Ints(pupils)

		res := pupils[n-1] - pupils[n]
		if res < 0 {
			res = -res
		}

		fmt.Println(res)
	}
}
