// https://codeforces.com/problemset/problem/672/B

package main

import "fmt"

func main() {
	var n int
	var s string
	fmt.Scan(&n, &s)

	if n > 26 {
		fmt.Println(-1)
	} else {
		ar := make([]int, 26)
		for i := range s {
			ar[s[i]-'a']++
		}

		used := 0
		for i := range ar {
			if ar[i] > 0 {
				used++
			}
		}

		res := n - used
		fmt.Println(res)
	}
}
