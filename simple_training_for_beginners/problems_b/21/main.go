// https://codeforces.com/problemset/problem/1150/B

package main

import "fmt"

type Grid [][]byte

func (g *Grid) PutFigure(i, j, n int) bool {
	if i+2 >= n || j-1 < 0 || j+1 >= n {
		return false
	}
	if (*g)[i+1][j] == '.' && (*g)[i+2][j] == '.' &&
		(*g)[i+1][j-1] == '.' && (*g)[i+1][j+1] == '.' {
		(*g)[i][j] = '#'
		(*g)[i+1][j] = '#'
		(*g)[i+2][j] = '#'
		(*g)[i+1][j-1] = '#'
		(*g)[i+1][j+1] = '#'
		return true
	}
	return false
}

func main() {
	var n int
	fmt.Scan(&n)

	grid := Grid(make([][]byte, n))
	s := ""
	for i := range grid {
		fmt.Scan(&s)
		grid[i] = []byte(s)
	}

	isPossibleToFill := true
	for i := 0; i < n; i++ {
		for j := range grid[i] {
			if grid[i][j] == '.' {
				isPossibleToFill = grid.PutFigure(i, j, n)
			}
			if !isPossibleToFill {
				i = n
				break
			}
		}
	}

	if isPossibleToFill {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
