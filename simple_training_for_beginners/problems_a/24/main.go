// https://codeforces.com/problemset/problem/92/A

package main

import "fmt"

func main() {
	var n, m int
	fmt.Scanf("%d %d\n", &n, &m)

	chipsInRound := (n*n + n) / 2
	m = m % chipsInRound

	for i := 1; i <= n; i++ {
		if (m - i) < 0 {
			break
		}
		m -= i
	}

	fmt.Println(m)
}
