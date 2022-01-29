// https://codeforces.com/problemset/problem/680/B

package main

import "fmt"

func main() {
	var n, a int
	fmt.Scan(&n, &a)
	a--

	cities := make([]int, n)
	for i := range cities {
		fmt.Scan(&cities[i])
	}

	res := 0
	l := a
	r := n - 1 - a
	min := l
	if l < r {
		for i := a + l + 1; i < n; i++ {
			res += cities[i]
		}
		min = l
	} else if l > r {
		for i := 0; i < a-r; i++ {
			res += cities[i]
		}
		min = r
	}

	for i := 0; i <= min; i++ {
		if cities[a+i]+cities[a-i] == 2 {
			res += 2
		}
	}
	res -= cities[a]

	fmt.Println(res)
}
