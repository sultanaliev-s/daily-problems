package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var n int
	fmt.Fscan(in, &n)

	costs := make([]int, n)
	for i := range costs {
		fmt.Fscan(in, &costs[i])
	}
	sort.Ints(costs)

	var days int
	fmt.Fscan(in, &days)
	for i := 0; i < days; i++ {
		var money int
		fmt.Fscan(in, &money)
		l, r, res := 0, n-1, -1
		for l <= r {
			m := (l + r) / 2
			if costs[m] <= money {
				l = m + 1
			} else {
				res = m
				r = m - 1
			}
		}

		if res == 0 {
			fmt.Fprintln(out, 0)
		} else if res == -1 {
			fmt.Fprintln(out, n)
		} else {
			fmt.Fprintln(out, res)
		}
	}
}
