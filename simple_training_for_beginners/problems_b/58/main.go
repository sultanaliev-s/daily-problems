// https://codeforces.com/problemset/problem/1051/B

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var l, r int64
	fmt.Scan(&l, &r)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	fmt.Fprintln(out, "YES")
	for i := l; i < r; i += 2 {
		fmt.Fprintln(out, i, i+1)
	}
}
