// https://codeforces.com/problemset/problem/1216/B

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(in, &n)

	targets := make([]struct {
		index      int
		durability int
	}, n)

	for i := range targets {
		fmt.Fscan(in, &targets[i].durability)
		targets[i].index = i + 1
	}

	sort.Slice(targets, func(i, j int) bool {
		return targets[i].durability > targets[j].durability
	})

	shotsNeeded := 0
	shotTargets := 0
	for i := range targets {
		shotsNeeded += targets[i].durability*shotTargets + 1
		shotTargets++
	}

	fmt.Println(shotsNeeded)
	for i := range targets {
		if i != 0 {
			fmt.Print(" ")
		}
		fmt.Print(targets[i].index)
	}
	fmt.Println()
}
