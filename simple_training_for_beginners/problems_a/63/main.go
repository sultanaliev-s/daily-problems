// https://codeforces.com/problemset/problem/669/A

package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	res := n / 3 * 2
	if n%3 != 0 {
		res++
	}

	fmt.Println(res)
}
