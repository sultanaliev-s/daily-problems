// https://codeforces.com/problemset/problem/859/B

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
			res = w*4 + 2
		} else {
			res = w*4 + 4
		}
	} else {
		res = w * 4
	}

	fmt.Println(res)
}
