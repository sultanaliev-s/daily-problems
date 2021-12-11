// https://codeforces.com/problemset/problem/136/A

package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d\n", &n)

	giftedBy := make([]int, n)
	var x int
	for i := 1; i <= n; i++ {
		fmt.Scanf("%d", &x)
		giftedBy[x-1] = i
	}

	for i, v := range giftedBy {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(v)
	}
	fmt.Println()
}
