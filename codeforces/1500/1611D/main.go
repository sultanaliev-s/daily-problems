package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var test int
	fmt.Fscan(in, &test)
	for ; test > 0; test-- {
		var n int
		fmt.Fscan(in, &n)

		ancestors := make([]int, n)
		for i := range ancestors {
			fmt.Fscan(in, &ancestors[i])
		}
		perm := make([]int, n)
		for i := range perm {
			fmt.Fscan(in, &perm[i])
		}

		root := 0
		for i := range ancestors {
			if i+1 == ancestors[i] {
				root = i
				break
			}
		}

		if root != perm[0]-1 {
			fmt.Fprintln(out, -1)
			continue
		}

		isPossible := true
		weight := 1
		weights := make([]int, n)
		ans := make([]int, n)
		for i := 1; i < n; i++ {
			p := perm[i] - 1
			a := ancestors[p] - 1
			w := weights[a]
			if w == 0 && a != root {
				isPossible = false
				break
			}
			ans[p] = weight - w
			weights[p] = ans[p] + w
			weight++
		}

		if !isPossible {
			fmt.Fprintln(out, -1)
		} else {
			for i := range ans {
				if i != 0 {
					fmt.Fprint(out, " ")
				}
				fmt.Fprint(out, ans[i])
			}
			fmt.Fprintln(out)
		}
	}
}
