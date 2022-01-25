// https://codeforces.com/problemset/problem/805/B

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	fmt.Scan(&n)

	out := bufio.NewWriter(os.Stdout)
	s := "aabb"
	for i := 0; i < n/4; i++ {
		fmt.Fprint(out, s)
	}

	fmt.Fprintln(out, s[:n%4])
	out.Flush()
}
