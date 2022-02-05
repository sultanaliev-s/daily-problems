// https://codeforces.com/problemset/problem/155/B

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Pair struct {
	p int
	e int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(in, &n)

	cards := make([]Pair, n)
	for i := range cards {
		fmt.Fscan(in, &cards[i].p, &cards[i].e)
	}

	sort.Slice(cards, func(i, j int) bool {
		if cards[i].e == cards[j].e {
			return cards[i].p > cards[j].p
		}
		return cards[i].e > cards[j].e
	})

	res := 0
	es := 1
	for i := 0; i < n && es > 0; i, es = i+1, es-1 {
		res += cards[i].p
		es += cards[i].e
	}

	fmt.Println(res)
}
