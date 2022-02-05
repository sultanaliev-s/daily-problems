// https://codeforces.com/problemset/problem/106/B

package main

import (
	"fmt"
	"sort"
)

type Comp struct {
	id    int
	speed int
	ram   int
	hdd   int
	cost  int
}

func main() {
	var n int
	fmt.Scan(&n)

	comps := make([]Comp, n)
	for i := range comps {
		fmt.Scan(&comps[i].speed, &comps[i].ram, &comps[i].hdd, &comps[i].cost)
		comps[i].id = i + 1
	}

	notOutdated := make([]Comp, 0, n)
	for i := range comps {
		isOutdated := false
		for j := range comps {
			if comps[i].speed < comps[j].speed &&
				comps[i].ram < comps[j].ram &&
				comps[i].hdd < comps[j].hdd {
				isOutdated = true
				break
			}
		}
		if !isOutdated {
			notOutdated = append(notOutdated, comps[i])
		}
	}

	sort.Slice(notOutdated, func(i, j int) bool {
		return notOutdated[i].cost < notOutdated[j].cost
	})

	fmt.Println(notOutdated[0].id)
}
