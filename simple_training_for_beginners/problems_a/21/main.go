// https://codeforces.com/problemset/problem/112/A

package main

import (
	"fmt"
	"strings"
)

func main() {
	var a, b string
	fmt.Scanln(&a)
	fmt.Scanln(&b)

	a = strings.ToLower(a)
	b = strings.ToLower(b)

	if a < b {
		fmt.Println(-1)
	} else if a > b {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}
