// https://codeforces.com/problemset/problem/899/C

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var n int64
	fmt.Scan(&n)

	gigaSum := (n * (n + 1)) / 2
	res := gigaSum % 2
	fmt.Fprintln(out, res)

	var s, f int64 // first, second
	q := n / 4
	p := n/2 - q
	x := n
	ar := make([]int64, 0, n/2+1)
	for i := int64(0); i < q; i++ {
		ar = append(ar, x)
		f += x
		s += x - 1
		x -= 2
	}
	for i := int64(0); i < p; i++ {
		ar = append(ar, x-1)
		s += x
		f += x - 1
		x -= 2
	}

	if x != 0 && f < s {
		ar = append(ar, 1)
	}

	fmt.Fprint(out, len(ar))
	for i := range ar {
		fmt.Fprint(out, " ", ar[i])
	}
	fmt.Fprintln(out)
}
