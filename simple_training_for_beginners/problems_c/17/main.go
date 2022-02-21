// https://codeforces.com/problemset/problem/1272/C

package main

import "fmt"

func main() {
	var n, k int
	var s string
	fmt.Scan(&n, &k, &s)

	avail := make(map[byte]bool)
	for i, x := 0, ""; i < k; i++ {
		fmt.Scan(&x)
		avail[x[0]] = true
	}

	var res int64
	for i := 0; i < n; i++ {
		if avail[s[i]] {
			j := i + 1
			for ; j < n && avail[s[j]]; j++ {
			}

			t := int64(j - i)
			res += t * (t + 1) / 2
			i = j - 1
		}
	}

	fmt.Println(res)
}
