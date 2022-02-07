// https://codeforces.com/problemset/problem/6/B

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n, m int
	var pColour string
	fmt.Fscan(in, &n, &m, &pColour)

	p := pColour[0]

	colours := make([][]byte, n)
	var s string
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &s)
		colours[i] = []byte(s)
	}

	coords := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	met := []byte{p, '.'}
	res := 0
	for i := range colours {
		for j := range colours[i] {
			if colours[i][j] == p {
				for k := range coords {
					tx, ty := i+coords[k][0], j+coords[k][1]
					if 0 <= tx && tx < n && 0 <= ty && ty < m &&
						!hasMet(met, colours[tx][ty]) {
						res++
						met = append(met, colours[tx][ty])
					}
				}
			}
		}
	}

	fmt.Println(res)
}

func hasMet(a []byte, c byte) bool {
	for i := range a {
		if a[i] == c {
			return true
		}
	}
	return false
}
