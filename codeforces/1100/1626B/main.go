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
	var test int
	fmt.Fscan(in, &test)
	for ; test > 0; test-- {
		var s string
		fmt.Fscan(in, &s)

		found := false
		for i := len(s) - 1; i > 1; i-- {
			if t := s[i] - '0' + s[i-1] - '0'; t >= 10 {
				fmt.Fprintf(out, "%s%d%s\n", s[:i-1], t, s[i+1:])
				found = true
				break
			}
		}

		if !found {
			t := s[0] - '0' + s[1] - '0'
			fmt.Fprintf(out, "%d%s\n", t, s[2:])
		}
	}
}
