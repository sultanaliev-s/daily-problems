// https://codeforces.com/problemset/problem/625/C

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

	var n, k int
	fmt.Fscan(in, &n, &k)

	ar := make([][]int, n)
	for i := range ar {
		ar[i] = make([]int, n)
	}

	s := 1
	for i := 0; i < n; i++ {
		for j := 0; j < k-1; j++ {
			ar[i][j] = s
			s++
		}
	}
	for i := 0; i < n; i++ {
		for j := k - 1; j < n; j++ {
			ar[i][j] = s
			s++
		}
	}

	sum := 0
	for i := 0; i < n; i++ {
		sum += ar[i][k-1]
	}

	fmt.Fprintln(out, sum)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j > 0 {
				fmt.Fprint(out, " ")
			}
			fmt.Fprint(out, ar[i][j])
		}
		fmt.Fprintln(out)
	}

}
