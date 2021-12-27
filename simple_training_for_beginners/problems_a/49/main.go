// https://codeforces.com/problemset/problem/723/A

package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scanf("%d %d %d\n", &a, &b, &c)

	length := Max(a, Max(b, c)) - Min(a, Min(b, c))

	fmt.Println(length)
}

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
