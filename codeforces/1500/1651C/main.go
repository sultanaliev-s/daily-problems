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
	var t int
	fmt.Fscan(in, &t)
	for ; t > 0; t-- {
		var n int
		fmt.Fscan(in, &n)

		a := make([]int64, n)
		b := make([]int64, n)
		for i := range a {
			fmt.Fscan(in, &a[i])
		}
		for i := range b {
			fmt.Fscan(in, &b[i])
		}

		m := n - 1
		cands1 := [3]int{0, m, bestCand(b, a[0])}
		cands2 := [3]int{0, m, bestCand(b, a[m])}

		var min int64 = 5_000_000_000
		for _, v := range cands1 {
			for _, u := range cands2 {
				cur := abs(a[0]-b[v]) + abs(a[m]-b[u])

				if v != 0 && u != 0 {
					cur += abs(b[0] - a[bestCand(a, b[0])])
				}
				if v != m && u != m {
					cur += abs(b[m] - a[bestCand(a, b[m])])
				}

				min = Min(cur, min)
			}
		}

		res := min
		fmt.Fprintln(out, res)
	}
}

func abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}
func Min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func bestCand(ar []int64, x int64) int {
	best := int64(5_000_000_000)
	pos := 0
	for i := range ar {
		if t := abs(x - ar[i]); t < best {
			best = t
			pos = i
		}
	}
	return pos
}
