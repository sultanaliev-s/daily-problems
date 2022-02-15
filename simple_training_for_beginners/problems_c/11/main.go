// https://codeforces.com/problemset/problem/1392/C

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
	var tests int
	fmt.Fscan(in, &tests)
	for ; tests > 0; tests-- {
		var n int
		fmt.Fscan(in, &n)

		ar := make([]int64, n)
		for i := range ar {
			fmt.Fscan(in, &ar[i])
		}

		var res int64
		for i := n - 1; i > 0; i-- {
			if ar[i-1] > ar[i] {
				res += ar[i-1] - ar[i]
			}
		}

		fmt.Fprintln(out, res)
	}
}
