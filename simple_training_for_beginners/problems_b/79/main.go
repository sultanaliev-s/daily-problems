// https://codeforces.com/problemset/problem/268/B

package main

import "fmt"

func main() {
	var n int64
	fmt.Scan(&n)

	res := int64(0)
	for i := n; i > 0; i-- {
		res += i + (n-i)*(i-1)
	}

	fmt.Println(res)
}
