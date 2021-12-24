// https://codeforces.com/problemset/problem/50/A

package main

import "fmt"

func main() {
	var m, n int
	fmt.Scanf("%d %d\n", &m, &n)

	dominos := m * n / 2

	fmt.Println(dominos)
}
