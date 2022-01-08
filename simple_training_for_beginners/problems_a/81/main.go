// https://codeforces.com/problemset/problem/263/A

package main

import "fmt"

func main() {
	var n int
	for i, x := 0, 0; i < 25; i++ {
		fmt.Scan(&x)
		if x == 1 {
			n = i
		}
	}

	r := n / 5
	c := n % 5
	res := abs(2-r) + abs(2-c)
	fmt.Println(res)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
