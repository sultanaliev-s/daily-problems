// https://codeforces.com/problemset/problem/747/C

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
	var n, q int
	fmt.Fscan(in, &n, &q)

	servers := make([]int, n)
	var sec, nServs, dur int
	for i := 0; i < q; i++ {
		fmt.Fscan(in, &sec, &nServs, &dur)
		free := 0
		for j := range servers {
			if servers[j] < sec {
				free++
			}
		}
		if nServs <= free {
			sum := 0
			for j, k := 0, 0; j < n && k < nServs; j++ {
				if servers[j] < sec {
					servers[j] = sec + dur - 1
					sum += j + 1
					k++
				}
			}
			fmt.Fprintln(out, sum)
		} else {
			fmt.Fprintln(out, -1)
		}
	}
}
