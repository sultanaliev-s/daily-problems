// https://codeforces.com/problemset/problem/158/A

package main

import "fmt"

func main() {
	var n, k int
	fmt.Scanf("%d %d\n", &n, &k)

	var nForNextRound int
	var x int
	for i := 0; i < k; i++ {
		fmt.Scanf("%d", &x)
		if i < k && x > 0 {
			nForNextRound++
		}
	}

	for i := 0; i < n-k; i++ {
		var y int
		fmt.Scanf("%d", &y)
		if y == x && y > 0 {
			nForNextRound++
		}
	}
	fmt.Scanln()

	fmt.Println(nForNextRound)
}
