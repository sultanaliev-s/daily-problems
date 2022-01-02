// https://codeforces.com/problemset/problem/629/A

package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	r, c := make([]int, n), make([]int, n)
	var line string
	res := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&line)
		for j := 0; j < n; j++ {
			if line[j] == 'C' {
				res += r[j] + c[i]
				r[j]++
				c[i]++
			}
		}
	}

	fmt.Println(res)
}
