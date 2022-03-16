// https://codeforces.com/problemset/problem/1271/C

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n, sx, sy int
	fmt.Fscan(in, &n, &sx, &sy)

	var x, y int
	var l, u, r, d, lu, ru, rd, ld int
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &x, &y)
		if x == sx {
			if y > sy {
				u++
			} else {
				d++
			}
		} else if y == sy {
			if x > sx {
				r++
			} else {
				l++
			}
		} else if x < sx && y > sy {
			lu++
		} else if x > sx && y > sy {
			ru++
		} else if x > sx && y < sy {
			rd++
		} else if x < sx && y < sy {
			ld++
		}
	}

	vars := [4]int{}
	vars[0] = l + lu + ld
	vars[1] = u + lu + ru
	vars[2] = r + ru + rd
	vars[3] = d + ld + rd

	max := 0
	for i := range vars {
		if vars[max] < vars[i] {
			max = i
		}
	}

	type pair struct {
		x int
		y int
	}
	ans := [8]pair{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}

	fmt.Println(vars[max])
	fmt.Println(sx+ans[max].x, sy+ans[max].y)
}
