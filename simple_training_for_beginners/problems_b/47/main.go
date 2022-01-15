// https://codeforces.com/problemset/problem/1199/B

package main

import "fmt"

func main() {
	var h, l float64
	fmt.Scan(&h, &l)

	ans := (l*l - h*h) / (2 * h)

	fmt.Println(ans)
}
