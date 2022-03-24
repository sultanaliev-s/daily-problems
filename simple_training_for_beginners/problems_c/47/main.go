// https://codeforces.com/problemset/problem/1391/C

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
	var n int64
	fmt.Fscan(in, &n)

	var mod, factorial, exclude int64
	mod = 1_000_000_007
	factorial = 1
	exclude = 1
	for i := int64(2); i <= n; i++ {
		factorial = (factorial * i) % mod
		exclude = (exclude * 2) % mod
	}

	res := (factorial - exclude) % mod
	if res < 0 {
		res += mod
	}

	fmt.Fprintln(out, res)
}
