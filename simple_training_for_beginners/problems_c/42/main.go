// https://codeforces.com/problemset/problem/1375/C

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

		br := make([]int, n)
		bi := 0
		br[bi] = ar[0]
		for i := 1; i < n; i++ {
			if br[bi] < ar[i] {
				for bi > 0 && br[bi-1] < ar[i] {
					bi--
				}
			} else {
				bi++
				br[bi] = ar[i]
			}
		}

		if bi == 0 {
			fmt.Fprintln(out, "YES")

		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}
