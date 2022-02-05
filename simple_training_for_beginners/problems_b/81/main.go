// https://codeforces.com/problemset/problem/157/B

package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	r := make([]int, n)
	for i := range r {
		fmt.Scan(&r[i])
	}
	sort.Sort(sort.Reverse(sort.IntSlice(r)))

	var res float64
	for i := 0; i < n-1; i += 2 {
		res += math.Pi*float64(r[i]*r[i]) - math.Pi*float64(r[i+1]*r[i+1])
	}
	if n%2 != 0 {
		res += math.Pi * float64(r[n-1]*r[n-1])
	}

	fmt.Println(res)
}
