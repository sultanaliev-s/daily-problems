// https://codeforces.com/problemset/problem/888/B

package main

import "fmt"

func main() {
	var n int
	var line string
	fmt.Scan(&n, &line)

	var u, d, r, l int
	for i := range line {
		switch line[i] {
		case 'U':
			u++
		case 'D':
			d++
		case 'R':
			r++
		case 'L':
			l++
		}
	}

	res := min(u, d)*2 + min(r, l)*2

	fmt.Println(res)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
