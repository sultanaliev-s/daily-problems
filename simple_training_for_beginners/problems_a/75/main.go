// https://codeforces.com/problemset/problem/509/A

package main

import "fmt"

func main() {
	var n int64
	fmt.Scan(&n)

	a := 2*n - 2
	b := n - 1

	res := mul(b+1, a) / mul(1, a-b)

	fmt.Println(res)
}

func mul(a, b int64) int64 {
	res := int64(1)
	for i := a; i <= b; i++ {
		res *= i
	}

	return res
}
