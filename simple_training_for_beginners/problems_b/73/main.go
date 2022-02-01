// https://codeforces.com/problemset/problem/620/B

package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	seg := [10]int{6, 2, 5, 5, 4, 5, 6, 3, 7, 6}
	res := 0
	for i, j := a, 0; i <= b; i++ {
		j = i
		for j != 0 {
			res += seg[j%10]
			j /= 10
		}
	}

	fmt.Println(res)
}
