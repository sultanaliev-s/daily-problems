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
		var n int
		var s string
		fmt.Fscan(in, &n, &s)

		beg := 0
		end := n/2 - 1
		found := false
		for end < n {
			prev := beg - 1
			next := end + 1
			if next < n && s[next] == '0' {
				fmt.Fprintln(out, beg+1, next+1, beg+1, end+1)
				found = true
				break
			} else if beg > 0 && s[prev] == '0' {
				fmt.Fprintln(out, prev+1, end+1, beg+1, end+1)
				found = true
				break
			}
			beg++
			end++
		}

		if found {
			continue
		}
		beg = 0
		end = n/2 - 1
	esc:
		for end < n {
			if equalsZero(s[beg : end+1]) {
				fmt.Fprintln(out, beg+1, end+1, 1, n)
				break
			}
			b := beg + 1
			f := end + 1
			for f < n {
				if equals(s[beg:end+1], s[b:f+1]) {
					fmt.Fprintln(out, beg+1, end+1, b+1, f+1)
					break esc
				}
				b++
				f++
			}
			beg++
			end++
		}
	}
}

func equals(s1, s2 string) bool {
	if len(s1) == len(s2) {
		for i := range s1 {
			if s1[i] != s2[i] {
				return false
			}
		}
		return true
	}
	return false
}

func equalsZero(s string) bool {
	for i := range s {
		if s[i] != '0' {
			return false
		}
	}
	return true
}
