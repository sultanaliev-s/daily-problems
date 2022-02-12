// https://codeforces.com/problemset/problem/1324/C

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
	var t int
	fmt.Fscan(in, &t)
	for ; t > 0; t-- {
		var s string
		fmt.Fscan(in, &s)

		l := 0
		max := -1
		for i := range s {
			if s[i] == 'L' {
				l++
			} else {
				if max < l {
					max = l
				}
				l = 0
			}
		}
		if max < l {
			max = l
		}

		if max == -1 {
			fmt.Fprintln(out, len(s)+1)
		} else {
			fmt.Fprintln(out, max+1)
		}
	}
}
