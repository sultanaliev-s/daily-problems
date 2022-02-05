// https://codeforces.com/problemset/problem/127/B

package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	sticks := make([]int, 101)
	for i, x := 0, 0; i < n; i++ {
		fmt.Scan(&x)
		sticks[x]++
	}

	pairs := 0
	for i := range sticks {
		pairs += sticks[i] / 2
	}

	frames := pairs / 2

	fmt.Println(frames)
}
