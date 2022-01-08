// https://codeforces.com/problemset/problem/381/A

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(in, &n)

	cards := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &cards[i])
	}

	var sereja, dima int
	left := 0
	right := n - 1
	isDimasTurn := false
	player := &dima
	for left <= right {
		if isDimasTurn {
			player = &dima
		} else {
			player = &sereja
		}
		isDimasTurn = !isDimasTurn

		if cards[left] > cards[right] {
			*player += cards[left]
			left++
		} else {
			*player += cards[right]
			right--
		}
	}

	fmt.Println(sereja, dima)
}
