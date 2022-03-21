// https://codeforces.com/problemset/problem/1312/C

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var t int
	fmt.Fscan(in, &t)
	for ; t > 0; t-- {
		var n, k int64
		fmt.Fscan(in, &n, &k)

		ar := make([]int64, n)
		for i := range ar {
			fmt.Fscan(in, &ar[i])
		}

		isPossible := true

		used := make([]bool, 62)
	out:
		for _, v := range ar {
			if v != 0 {
				basa := f(v, k)
				for i := range basa {
					if basa[i] == '1' {
						if used[i] {
							isPossible = false
							break out
						} else {
							used[i] = true
						}
					} else if basa[i] != '0' {
						isPossible = false
						break out
					}
				}
			}
		}

		if isPossible {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}

func f(v, k int64) string {
	var s string
	var t int64
	for {
		t = v % k
		if t > 1 {
			t = 2
		}
		s += strconv.FormatInt(t, 10)
		v /= k
		if v == 0 {
			break
		}
	}

	return s
}
