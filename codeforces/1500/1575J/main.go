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
	var n, m, k int
	fmt.Fscan(in, &n, &m, &k)

	grid := make([][]int, n)
	for i := range grid {
		grid[i] = make([]int, m)
		for j := range grid[i] {
			fmt.Fscan(in, &grid[i][j])
		}
	}

	for i := 0; i < k; i++ {
		var x, y int
		fmt.Fscan(in, &x)
		x--

		for y != n {
			cmd := grid[y][x]
			grid[y][x] = 2
			if cmd == 1 {
				x++
			} else if cmd == 2 {
				y++
			} else {
				x--
			}
		}

		if i != 0 {
			fmt.Fprint(out, " ")
		}
		fmt.Fprint(out, x+1)
	}
	fmt.Fprintln(out)
}
