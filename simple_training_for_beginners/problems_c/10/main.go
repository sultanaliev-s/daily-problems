// https://codeforces.com/problemset/problem/1399/C

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

		ar := make([]int, n)
		for i := range ar {
			fmt.Fscan(in, &ar[i])
		}

		ress := make([]int, 0, 100)
		for s := 2; s < 2*n+1; s++ {
			res := 0
			b := make([]bool, n)
			for i := range ar {
				for j := range ar {
					if i != j && !b[i] && !b[j] && ar[i]+ar[j] == s {
						res++
						b[i] = true
						b[j] = true
					}
				}
			}

			ress = append(ress, res)
		}

		max := 0
		for i := range ress {
			if ress[i] > ress[max] {
				max = i
			}
		}

		fmt.Fprintln(out, ress[max])
	}
}
