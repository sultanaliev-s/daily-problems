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
	var tests int
	fmt.Fscan(in, &tests)
	for test := 0; test < tests; test++ {
		var s, t string
		fmt.Fscan(in, &s, &t)

		if len(t) > len(s) {
			fmt.Fprintln(out, "NO")
			continue
		}

		n, m := len(s), len(t)
		ar := []byte(s)
		matched := 0
		for i, j := n-1, m-1; i >= 0 && j >= 0; {
			if ar[i] != t[j] {
				ar[i] = '*'
				if i > 0 {
					ar[i-1] = '*'
				}
				i -= 2
			} else {
				i--
				j--
				matched++
			}
		}

		isPossible := matched == m

		if isPossible {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}
