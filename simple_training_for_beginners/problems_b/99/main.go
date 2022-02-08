// https://codeforces.com/problemset/problem/1099/B

package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	w := int(math.Sqrt(float64(n)))
	res := 0
	if sq := w * w; sq < n {
		if n-sq <= w {
			res = w*2 + 1
		} else {
			res = w*2 + 2
		}
	} else {
		res = w * 2
	}

	fmt.Println(res)
}
