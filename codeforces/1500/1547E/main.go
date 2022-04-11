package main

import (
	"bufio"
	"fmt"
	"os"
)

const MAX_INT int = 2_000_000_000

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var test int
	fmt.Fscan(in, &test)
	for ; test > 0; test-- {
		var n, k int
		fmt.Fscan(in, &n, &k)

		a := make([]int, k)
		t := make([]int, k)
		ar := make([]int, n)
		for i := range a {
			fmt.Fscan(in, &a[i])
		}
		for i := range t {
			fmt.Fscan(in, &t[i])
		}
		for i := range a {
			ar[a[i]-1] = t[i]
		}
		for i := range ar {
			if ar[i] == 0 {
				ar[i] = MAX_INT
			}
		}

		L := make([]int, n)
		R := make([]int, n)
		p := MAX_INT - 1
		for i := 0; i < n; i++ {
			p = Min(p+1, ar[i])
			L[i] = p
		}
		p = MAX_INT - 1
		for i := n - 1; i >= 0; i-- {
			p = Min(p+1, ar[i])
			R[i] = p
		}

		for i := range L {
			if i != 0 {
				fmt.Fprint(out, " ")
			}
			fmt.Fprint(out, Min(L[i], R[i]))
		}
		fmt.Fprintln(out)
	}
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
