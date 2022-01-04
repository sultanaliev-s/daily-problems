// https://codeforces.com/problemset/problem/581/A

package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	if a > b {
		a, b = b, a
	}

	diffColors := a
	sameColor := (b - a) / 2

	fmt.Println(diffColors, sameColor)
}
