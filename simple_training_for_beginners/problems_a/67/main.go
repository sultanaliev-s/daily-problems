// https://codeforces.com/problemset/problem/617/A

package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	res := n / 5
	if n%5 != 0 {
		res++
	}

	fmt.Println(res)
}
