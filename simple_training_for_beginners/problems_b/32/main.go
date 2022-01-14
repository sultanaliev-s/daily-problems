// https://codeforces.com/problemset/problem/801/B

package main

import (
	"bufio"
	"fmt"
	"os"
)

func f(x, y string) string {
	z := make([]byte, len(x))
	for i := range x {
		z[i] = min(x[i], y[i])
	}
	return string(z)
}

func min(a, b byte) byte {
	if a < b {
		return a
	}
	return b
}

func main() {
	in := bufio.NewReader(os.Stdin)
	var x, y string
	fmt.Fscan(in, &x, &y)

	z := f(x, y)

	if z != y {
		fmt.Println(-1)
	} else {
		fmt.Println(z)
	}
}
