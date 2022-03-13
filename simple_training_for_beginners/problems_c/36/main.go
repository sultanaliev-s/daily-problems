// https://codeforces.com/problemset/problem/1367/C

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var t int
	fmt.Scan(&t)
	for ; t > 0; t-- {
		var n, k int
		var s string
		fmt.Scan(&n, &k, &s)

		ar := []byte(s)
		res := 0
		zero := k
		for i := 0; i < n; i++ {
			if ar[i] == '0' {
				zero++
				if zero > k {
					one := i
					for j := i + 1; j <= i+k && j < n; j++ {
						if ar[j] == '1' {
							one = j
						}
					}
					if one == i {
						res++
						zero = 0
					} else {
						i = one
						zero = 0
					}
				}
			} else {
				zero = 0
			}
		}

		fmt.Fprintln(out, res)
	}
}
