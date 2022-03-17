// https://codeforces.com/problemset/problem/1194/C

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
		var s, t, p string
		fmt.Fscan(in, &s, &t, &p)

		if len(s) > len(t) || len(s)+len(p) < len(t) {
			fmt.Fprintln(out, "NO")
			continue
		}

		j := 0
		for i := 0; i < len(t) && j < len(s); i++ {
			if t[i] == s[j] {
				j++
			}
		}

		if j != len(s) {
			fmt.Fprintln(out, "NO")
			continue
		}

		alphT := make([]int, 26)
		alphSP := make([]int, 26)
		for i := range t {
			alphT[t[i]-'a']++
		}
		for i := range s {
			alphSP[s[i]-'a']++
		}
		for i := range p {
			alphSP[p[i]-'a']++
		}

		isPossible := true
		for i := range alphSP {
			if alphT[i] > alphSP[i] {
				isPossible = false
			}
		}

		if isPossible {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}
