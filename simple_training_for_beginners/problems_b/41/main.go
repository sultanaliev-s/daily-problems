// https://codeforces.com/problemset/problem/1293/B

package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var res float64
	for ; n > 0; n-- {
		res += 1.0 / float64(n)
	}

	fmt.Println(res)
}
