// https://codeforces.com/problemset/problem/630/C

package main

import "fmt"

func main() {
	var n int64
	fmt.Scan(&n)

	var res, i int64
	for i = 1; i <= n; i++ {
		var p, j int64
		p = 1
		for j = 1; j <= i; j++ {
			p *= 2
		}
		res += p
	}

	fmt.Println(res)
}
