// https://codeforces.com/problemset/problem/688/B

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
	var s string
	fmt.Fscan(in, &s)
	fmt.Fprint(out, s)
	out.Flush()

	for i := len(s) - 1; i >= 0; i-- {
		fmt.Fprintf(out, "%c", s[i])
	}
	fmt.Fprintln(out)

}
