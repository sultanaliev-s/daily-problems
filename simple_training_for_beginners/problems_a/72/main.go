// https://codeforces.com/problemset/problem/546/A

package main

import "fmt"

func main() {
	var k, n, w int
	fmt.Scan(&k, &n, &w)

	needed := (w * (w + 1)) / 2 * k
	borrowing := needed - n
	if borrowing < 0 {
		borrowing = 0
	}

	fmt.Println(borrowing)
}
