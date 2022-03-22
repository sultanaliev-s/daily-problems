// https://codeforces.com/problemset/problem/1270/C

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

		ar := make([]int64, n)
		var sum, xor int64
		for i := range ar {
			fmt.Fscan(in, &ar[i])
			sum += ar[i]
			xor ^= ar[i]
		}

		if sum == 2*xor {
			fmt.Fprintln(out, 0)
			fmt.Fprintln(out)
		} else {
			sum2 := sum + xor
			fmt.Fprintln(out, 2)
			fmt.Fprintln(out, xor, sum2)
		}
	}
}
