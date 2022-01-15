// https://codeforces.com/problemset/problem/1223/B

package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)

	for ; t > 0; t-- {
		var s, t string
		fmt.Scan(&s, &t)

		found := false
		for i := 0; i < len(s); i++ {
			for j := 0; j < len(t); j++ {
				if s[i] == t[j] {
					found = true
					i = len(s)
					break
				}
			}
		}

		if found {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
