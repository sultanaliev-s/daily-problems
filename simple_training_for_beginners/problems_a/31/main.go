// https://codeforces.com/problemset/problem/47/A

package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scanf("%d\n", &n)

	a := math.Sqrt(float64(n * 2))
	b := int(a)
	c := b * (b + 1) / 2

	if c == n {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
