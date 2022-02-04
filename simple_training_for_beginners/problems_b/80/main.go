// https://codeforces.com/problemset/problem/177/B1

package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	res := n
	for n > 1 {
		for i := 2; i <= n; i++ {
			if n%i == 0 {
				res += n / i
				n /= i
				break
			}
		}
	}

	fmt.Println(res)
}
