// https://codeforces.com/problemset/problem/750/A

package main

import "fmt"

func main() {
	var n, k int
	fmt.Scanf("%d %d\n", &n, &k)

	mins := 240 - k
	problems := 0

	for i := 1; i <= n; i++ {
		if mins < i*5 {
			break
		}
		problems++
		mins -= i * 5
	}

	fmt.Println(problems)
}
