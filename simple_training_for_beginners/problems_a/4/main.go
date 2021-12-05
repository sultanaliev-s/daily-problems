// https://codeforces.com/problemset/problem/233/A

package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d\n", &n)

	if (n % 2) == 0 {
		for i := 1; i <= n; i += 2 {
			if i > 1 {
				fmt.Print(" ")
			}
			fmt.Print(i+1, i)
		}
		fmt.Println()
	} else {
		fmt.Println(-1)
	}
}
