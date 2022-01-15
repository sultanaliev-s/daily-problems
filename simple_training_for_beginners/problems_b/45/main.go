// https://codeforces.com/problemset/problem/1230/B

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
	var line string
	fmt.Fscan(in, &n, &k, &line)

	num := []byte(line)

	if n == 1 && k > 0 {
		num[0] = '0'
	} else if n > 1 && k > 0 && num[0] != '1' {
		num[0] = '1'
		k--
	}

	for i := 1; i < n; i++ {
		if num[i] != '0' && k > 0 {
			num[i] = '0'
			k--
		}
	}

	fmt.Fprintln(out, string(num))
}
