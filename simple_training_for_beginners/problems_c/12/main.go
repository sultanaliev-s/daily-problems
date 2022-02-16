// https://codeforces.com/problemset/problem/1385/C

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

		isFalling := false
		ind := 0
		for i := n - 1; i > 0; i-- {
			if !isFalling && ar[i-1] < ar[i] {
				isFalling = true
			} else if isFalling && ar[i-1] > ar[i] {
				ind = i
				break
			}
		}

		fmt.Fprintln(out, ind)
	}
}
