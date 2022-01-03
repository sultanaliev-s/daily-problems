// https://codeforces.com/problemset/problem/599/A

package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	if a > b {
		a, b = b, a
	}

	res := a
	if a+b < c {
		res += a + b
	} else {
		res += c
	}

	if a+c < b {
		res += a + c
	} else {
		res += b
	}

	fmt.Println(res)
}
