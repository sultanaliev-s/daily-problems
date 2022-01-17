// https://codeforces.com/problemset/problem/1117/B

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	var n int
	var m, k int64
	fmt.Fscan(in, &n, &m, &k)

	ems := make([]int64, n)
	for i := range ems {
		fmt.Fscan(in, &ems[i])
	}

	sort.Slice(ems, func(i, j int) bool {
		return ems[i] > ems[j]
	})

	fullCycles := m / (k + 1)
	remainder := m % (k + 1)

	res := fullCycles*k*ems[0] + fullCycles*ems[1] + remainder*ems[0]

	fmt.Fprintln(out, res)
	out.Flush()
}
