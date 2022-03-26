// https://codeforces.com/problemset/problem/1372/C

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

		end1 := 0
		for ; end1 < n; end1++ {
			if ar[end1] != end1+1 {
				break
			}
		}
		end2 := n - 1
		for ; end2 >= 0; end2-- {
			if ar[end2] != end2+1 {
				break
			}
		}

		isDerangement := true
		for i := end1; i <= end2; i++ {
			if ar[i] == i+1 {
				isDerangement = false
				break
			}
		}

		if end1 == n {
			fmt.Fprintln(out, 0)
		} else if isDerangement {
			fmt.Fprintln(out, 1)
		} else {
			fmt.Fprintln(out, 2)
		}

	}
}
