// https://codeforces.com/problemset/problem/734/B

package main

import "fmt"

func main() {
	var k2, k3, k5, k6 int
	fmt.Scan(&k2, &k3, &k5, &k6)

	lowest := min(k2, min(k5, k6))
	sum := lowest * 256
	k2 -= lowest
	sum += 32 * min(k2, k3)

	fmt.Println(sum)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
