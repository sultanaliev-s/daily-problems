// https://codeforces.com/problemset/problem/1360/C

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

		a := make([]int, n)
		for i := range a {
			fmt.Fscan(in, &a[i])
		}

		sort.Ints(a)
		even := 0
		for i := range a {
			if a[i]%2 == 0 {
				even++
			}
		}
		sequential := 0
		for i := 1; i < n; i++ {
			if a[i-1]+1 == a[i] {
				sequential++
				if i+2 < n {
					i++
				}
			}
		}

		if even%2 != 0 && sequential == 0 {
			fmt.Fprintln(out, "NO")
		} else {
			fmt.Fprintln(out, "YES")
		}
	}
}
