// https://codeforces.com/problemset/problem/1388/B

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	var t int
	fmt.Fscan(in, &t)

	for ; t > 0; t-- {
		var n int
		fmt.Fscan(in, &n)

		r := n*4 - n
		nines := r / 4

		for i := 0; i < nines; i++ {
			fmt.Fprint(out, "9")
		}

		eights := (n + 3) / 4
		for i := 0; i < eights; i++ {
			fmt.Fprint(out, "8")
		}
		fmt.Fprintln(out)
		out.Flush()
	}
}
