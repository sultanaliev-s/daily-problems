// https://codeforces.com/problemset/problem/1315/C

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

		ar := make([]int, n*2)
		taken := make([]bool, n*2+1)
		for i := 0; i < n*2; i += 2 {
			fmt.Fscan(in, &ar[i])
			taken[ar[i]] = true
		}

		isPossible := true
		for i := 0; i < n*2; i += 2 {
			hasBigger := false
			j := ar[i] + 1
			for ; j <= n*2; j++ {
				if !taken[j] {
					hasBigger = true
					break
				}
			}
			if !hasBigger {
				isPossible = false
				break
			}
			ar[i+1] = j
			taken[j] = true
		}

		if !isPossible {
			fmt.Fprintln(out, "-1")
		} else {
			for i := range ar {
				if i > 0 {
					fmt.Fprint(out, " ")
				}
				fmt.Fprint(out, ar[i])
			}
			fmt.Fprintln(out)
		}

	}
}
