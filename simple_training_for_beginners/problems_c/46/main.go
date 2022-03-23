// https://codeforces.com/problemset/problem/1400/C

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
		var s string
		var x int
		fmt.Fscan(in, &s, &x)
		n := len(s)

		ar := make([]int, n)
		for i := range ar {
			ar[i] = 1
		}

		for i := 0; i < n; i++ {
			if s[i] == '0' {
				ind1 := i - x
				ind2 := i + x
				if ind1 >= 0 {
					ar[ind1] = 0
				}
				if ind2 < n {
					ar[ind2] = 0
				}
			}
		}

		isPossible := true
		for i := range s {
			if s[i] == '1' {
				ok := false

				ind1 := i - x
				ind2 := i + x
				if ind1 >= 0 && ar[ind1] == 1 {
					ok = true
				}
				if ind2 < n && ar[ind2] == 1 {
					ok = true
				}
				if !ok {
					isPossible = false
					break
				}
			}
		}

		if isPossible {
			for i := range ar {
				if ar[i] == 0 {
					fmt.Fprint(out, 0)
				} else if ar[i] == 1 {
					fmt.Fprint(out, 1)
				} else {
					fmt.Fprint(out, 2)
				}
			}
			fmt.Fprintln(out)
		} else {
			fmt.Fprintln(out, -1)
		}
	}
}
