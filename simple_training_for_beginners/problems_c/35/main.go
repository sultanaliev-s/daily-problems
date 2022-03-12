// https://codeforces.com/problemset/problem/1401/C

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

		ar := make([]int, n)
		for i := range ar {
			fmt.Fscan(in, &ar[i])
		}

		min := ar[0]
		for i := range ar {
			if min > ar[i] {
				min = ar[i]
			}
		}

		isPossible := true
		for i := range ar {
			if gcd(ar[i], min) != min {
				isPossible = false
			}
		}

		if isPossible == false {
			isPossible = true
			for i := range ar {
				m := i
				for j := i; j < n; j++ {
					if ar[j] < ar[m] {
						m = j
					}
				}
				if m == i {
					continue
				}
				if gcd(ar[m], ar[i]) == min ||
					gcd(ar[m], min) == min && gcd(min, ar[i]) == min {
					ar[m], ar[i] = ar[i], ar[m]
				} else {
					isPossible = false
					break
				}
			}
		}

		if isPossible || sort.IsSorted(sort.IntSlice(ar)) {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}

func gcd(x, y int) int {
	if y == 0 {
		return x
	}
	return gcd(y, x%y)
}
