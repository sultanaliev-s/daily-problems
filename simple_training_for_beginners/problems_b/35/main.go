// https://codeforces.com/problemset/problem/16/B

package main

import (
	"fmt"
	"sort"
)

type Boxes struct {
	amount  int
	matches int
}

func main() {
	var n, nConts int
	fmt.Scan(&n, &nConts)

	conts := make([]Boxes, nConts)
	for i := range conts {
		fmt.Scan(&conts[i].amount, &conts[i].matches)
	}

	sort.Slice(conts, func(i, j int) bool {
		return conts[i].matches > conts[j].matches
	})

	matches := 0
	for i := range conts {
		t := n - conts[i].amount
		if t >= 0 {
			matches += conts[i].amount * conts[i].matches
			n = t
		} else if n > 0 {
			matches += n * conts[i].matches
			n = 0
		} else {
			break
		}
	}

	fmt.Println(matches)
}
