// https://codeforces.com/problemset/problem/701/A

package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	cards := make([][]int, 101)
	var x int
	max := 0
	min := 100
	for i := 1; i <= n; i++ {
		fmt.Scan(&x)
		cards[x] = append(cards[x], i)
		if max < x {
			max = x
		}
		if min > x {
			min = x
		}
	}

	sumOfTwo := min + max
	needed := 0
	a, b := 0, 0
	for i := range cards {
		for len(cards[i]) > 0 {
			needed = sumOfTwo - i
			a = cards[i][len(cards[i])-1]
			cards[i] = cards[i][:len(cards[i])-1]

			b = cards[needed][len(cards[needed])-1]
			cards[needed] = cards[needed][:len(cards[needed])-1]

			fmt.Println(a, b)
		}
	}
}
