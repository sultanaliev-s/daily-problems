// https://codeforces.com/problemset/problem/200/B

package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var sum, x int
	for i := 0; i < n; i++ {
		fmt.Scan(&x)
		sum += x
	}

	ratio := float64(sum) / float64(n)
	fmt.Println(ratio)
}
