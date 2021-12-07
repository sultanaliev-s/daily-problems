// https://codeforces.com/problemset/problem/151/A

package main

import "fmt"

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	var n, k, l, c, d, p, nl, np int
	fmt.Scanf("%d %d %d %d %d %d %d %d\n",
		&n, &k, &l, &c, &d, &p, &nl, &np)

	drink := k * l / nl
	limes := c * d
	salt := p / np
	result := min(drink, min(limes, salt)) / n
	fmt.Println(result)
}
