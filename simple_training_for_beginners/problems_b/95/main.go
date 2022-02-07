// https://codeforces.com/problemset/problem/1334/B

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
		var n, m int64
		fmt.Fscan(in, &n, &m)

		ar := make([]int64, n)
		for i := range ar {
			fmt.Fscan(in, &ar[i])
		}

		sort.Slice(ar, func(i, j int) bool {
			return ar[i] > ar[j]
		})

		var sum, c int64
		for i := int64(0); i < n; i++ {
			sum += ar[i]
			if sum/(c+1) < m {
				break
			}
			c++
		}

		fmt.Println(c)
	}
}
