// https://codeforces.com/problemset/problem/1380/C

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
	var t int
	fmt.Fscan(in, &t)
	for ; t > 0; t-- {
		var n int
		var lim int64
		fmt.Fscan(in, &n, &lim)

		ar := make([]int64, n)
		for i := range ar {
			fmt.Fscan(in, &ar[i])
		}

		sort.Slice(ar, func(i, j int) bool { return ar[i] > ar[j] })

		teams := 0
		var mems int64
		for i := 0; i < n; i++ {
			mems++
			if mems*ar[i] >= lim {
				teams++
				mems = 0
			}
		}

		fmt.Fprintln(out, teams)
	}
}
