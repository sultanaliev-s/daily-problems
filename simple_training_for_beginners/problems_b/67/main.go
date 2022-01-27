// https://codeforces.com/problemset/problem/732/B

package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	additionalWalks := 0
	days := make([]int, n)
	fmt.Scan(&days[0])
	for i := 1; i < n; i++ {
		fmt.Scan(&days[i])
		if sum := days[i] + days[i-1]; sum < k {
			days[i] += k - sum
			additionalWalks += k - sum
		}
	}

	fmt.Println(additionalWalks)
	for i := 0; i < n; i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(days[i])
	}
	fmt.Println()
}
