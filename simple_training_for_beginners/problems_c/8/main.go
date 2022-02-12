// https://codeforces.com/problemset/problem/1231/C

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n, m int
	fmt.Fscan(in, &n, &m)

	a := make([][]int, n)
	for i := range a {
		a[i] = make([]int, m)
		for j := range a[i] {
			fmt.Fscan(in, &a[i][j])
		}
	}

	isPossible := true
	sum := 0
	for i := n - 2; i > 0 && isPossible; i-- {
		for j := m - 2; j > 0; j-- {
			t := a[i][j]
			if a[i][j] == 0 {
				t = min(a[i+1][j]-1, a[i][j+1]-1)
			}
			if t > a[i-1][j] && t > a[i][j-1] &&
				t < a[i+1][j] && t < a[i][j+1] {
				a[i][j] = t
				sum += t
			} else {
				isPossible = false
				break
			}
		}
	}

	if isPossible {
		for i := 1; i < n; i++ {
			sum += a[i][0] + a[i][m-1]
			if a[i-1][0] >= a[i][0] || a[i-1][m-1] >= a[i][m-1] {
				isPossible = false
				goto L1
			}
		}
		for i := 1; i < m; i++ {
			sum += a[0][i] + a[n-1][i]
			if a[0][i-1] >= a[0][i] || a[n-1][i-1] >= a[n-1][i] {
				isPossible = false
				goto L1
			}
		}
		sum += a[0][0] - a[n-1][m-1]
	}

L1:

	if isPossible {
		fmt.Println(sum)
	} else {
		fmt.Println(-1)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
