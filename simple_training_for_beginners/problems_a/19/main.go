// https://codeforces.com/problemset/problem/119/A

package main

import "fmt"

func gcd(a, b int) int {
	if a == 0 {
		return b
	}

	return gcd(b%a, a)
}

func main() {
	var a, b, n int
	fmt.Scanf("%d %d %d\n", &a, &b, &n)

	var semenWins bool
	var playerNumbers = []int{a, b}
	for i := 0; ; i++ {
		x := gcd(playerNumbers[i%2], n)
		if x > n {
			semenWins = (i % 2) != 0
			break
		}
		n -= x
	}

	if semenWins {
		fmt.Println(0)
	} else {
		fmt.Println(1)
	}
}
