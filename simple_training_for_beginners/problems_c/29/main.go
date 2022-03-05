// https://codeforces.com/problemset/problem/919/C

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n, m, k int
	fmt.Fscan(in, &n, &m, &k)

	ar := make([]string, n)
	for i := range ar {
		fmt.Fscan(in, &ar[i])
	}

	res := 0
	for i := 0; i < n; i++ {
		cons := 0
		for j := 0; j < m; j++ {
			if ar[i][j] == '.' {
				cons++
			} else {
				if t := cons - k; t >= 0 {
					res += t + 1
				}
				cons = 0
			}
		}
		if t := cons - k; cons != 0 && t >= 0 {
			res += t + 1
		}
	}

	if k != 1 {
		for i := 0; i < m; i++ {
			cons := 0
			for j := 0; j < n; j++ {
				if ar[j][i] == '.' {
					cons++
				} else {
					if t := cons - k; t >= 0 {
						res += t + 1
					}
					cons = 0
				}
			}
			if t := cons - k; cons != 0 && t >= 0 {
				res += t + 1
			}
		}
	}

	fmt.Println(res)
}
