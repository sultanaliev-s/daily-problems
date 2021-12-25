// https://codeforces.com/problemset/problem/38/A

package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d\n", &n)

	years := make([]int, n-1)
	for i := 0; i < n-1; i++ {
		fmt.Scanf("%d", &years[i])
	}
	fmt.Scanln()

	var a, b int
	fmt.Scanf("%d %d\n", &a, &b)

	var yearsToPromote int
	for i := a; i < b; i++ {
		yearsToPromote += years[i-1]
	}

	fmt.Println(yearsToPromote)
}
