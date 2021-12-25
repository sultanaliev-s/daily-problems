// https://codeforces.com/problemset/problem/32/A

package main

import "fmt"

func main() {
	var n, d int
	fmt.Scanf("%d %d\n", &n, &d)

	soldiers := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &soldiers[i])
	}

	possibleWays := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i != j && Abs(soldiers[i]-soldiers[j]) <= d {
				possibleWays++
			}
		}
	}

	fmt.Println(possibleWays)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
