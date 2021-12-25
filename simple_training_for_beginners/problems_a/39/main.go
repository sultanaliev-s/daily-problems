// https://codeforces.com/problemset/problem/764/A

package main

import "fmt"

func main() {
	var n, m, z int
	fmt.Scanf("%d %d %d\n", &n, &m, &z)

	artistsToDrown := z / LCM(n, m)

	fmt.Println(artistsToDrown)
}

func LCM(a, b int) int {
	return a * b / GCD(a, b)
}

func GCD(a, b int) int {
	if b == 0 {
		return a
	}

	return GCD(b, a%b)
}
