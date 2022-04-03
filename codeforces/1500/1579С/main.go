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
	dots := ".................................................................."
	var test int
	fmt.Fscan(in, &test)
	for ; test > 0; test-- {
		var n, m, kk int
		fmt.Fscan(in, &n, &m, &kk)

		grid := make([][]byte, n)
		used := make([][]byte, n)
		var s string
		for i := range grid {
			fmt.Fscan(in, &s)
			grid[i] = []byte(s)
			used[i] = []byte(dots[:m])
		}

		for i := range grid {
			for j := range grid[i] {
				if grid[i][j] == '*' && isRoot(i, j, grid) {
					c := 0
					for k := 1; k <= i; k++ {
						left := j - k
						right := j + k
						if left >= 0 && right < m &&
							grid[i-k][left] == '*' && grid[i-k][right] == '*' {
							c++
						} else {
							break
						}
					}
					if c >= kk {
						for k := 1; k <= c; k++ {
							left := j - k
							right := j + k
							used[i-k][left] = '*'
							used[i-k][right] = '*'
						}
						used[i][j] = '*'
					}
				}
			}
		}

		isPossible := true
		for i := range used {
			for j := range used[i] {
				if grid[i][j] != used[i][j] {
					isPossible = false
				}
			}
		}
		if isPossible {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}

func isRoot(i, j int, grid [][]byte) bool {
	if i == 0 || j == 0 || j == len(grid[0])-1 {
		return false
	}
	return grid[i-1][j-1] == '*' && grid[i-1][j+1] == '*'
}
