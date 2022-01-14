// https://codeforces.com/problemset/problem/1359/B

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var t int
	fmt.Fscan(in, &t)

	for ; t > 0; t-- {
		var n, m, x, y int
		fmt.Fscan(in, &n, &m, &x, &y)

		var line string
		var price, px, py int
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &line)

			for j := 0; j < m; j++ {
				if line[j] == '.' {
					px += x
				}
			}

			for j := 0; j < m; j++ {
				if line[j] == '.' {
					if j+1 < m && line[j+1] == '.' {
						py += y
						j++
					} else {
						py += x
					}
				}
			}

			price += min(px, py)
			px, py = 0, 0
		}

		fmt.Println(price)
	}

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
