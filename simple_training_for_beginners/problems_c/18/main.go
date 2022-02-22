// https://codeforces.com/problemset/problem/1257/C

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
		fmt.Fscan(in, &n)

		ar := make([][]int, n+1)
		var x int
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &x)
			ar[x] = append(ar[x], i)
		}

		sort.Slice(ar, func(i, j int) bool {
			return len(ar[i]) > len(ar[j])
		})

		isPossible := n > 1 && len(ar[0]) > 1
		if !isPossible {
			fmt.Fprintln(out, -1)
		} else {
			dists := make([]int, 0, 100)
			for i := range ar {
				for j := 1; j < len(ar[i]); j++ {
					dists = append(dists, ar[i][j]-ar[i][j-1]+1)
				}
			}

			sort.Ints(dists)

			fmt.Fprintln(out, dists[0])
		}
	}

}
