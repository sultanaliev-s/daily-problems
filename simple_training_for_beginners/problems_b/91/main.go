// https://codeforces.com/problemset/problem/120/B

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fi, _ := os.Open("input.txt")
	defer fi.Close()
	in := bufio.NewReader(fi)
	fo, _ := os.Create("output.txt")
	defer fo.Close()
	out := bufio.NewWriter(fo)
	defer out.Flush()

	var n, k int
	fmt.Fscan(in, &n, &k)

	s := make([]int, n)
	for i := range s {
		fmt.Fscan(in, &s[i])
	}

	res := 0
	for i := k - 1; i < n; {
		if s[i] == 1 {
			res = i
			break
		}
		if i == n-1 {
			i = -1
		}
		i++
	}
	res++

	fmt.Fprintln(out, res)
}
