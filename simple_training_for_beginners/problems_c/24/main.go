// https://codeforces.com/problemset/problem/1102/C

package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, x, y int
	fmt.Scan(&n, &x, &y)

	ar := make([]int, n)
	for i := range ar {
		fmt.Scan(&ar[i])
	}

	if x > y {
		fmt.Println(n)
	} else {
		sort.Ints(ar)
		count := 0
		for i := range ar {
			if ar[i] > x {
				break
			}
			count++
		}
		res := (count + 1) / 2
		fmt.Println(res)
	}
}
