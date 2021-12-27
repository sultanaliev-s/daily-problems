// https://codeforces.com/problemset/problem/746/A

package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scanf("%d\n%d\n%d\n", &a, &b, &c)

	nCompotesLemons := a
	nCompotesApples := b / 2
	nCompotesPears := c / 4

	compotes := Min(nCompotesLemons, (Min(nCompotesApples, nCompotesPears)))
	fruits := compotes + compotes*2 + compotes*4

	fmt.Println(fruits)
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
