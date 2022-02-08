// https://codeforces.com/problemset/problem/1146/B

package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	a := 0
	ai := 0
	for i := range s {
		if s[i] == 'a' {
			a++
			ai = i
		}
	}
	n := len(s)
	b := n - (n-a+1)/2
	isPossible := true
	for i, j := 0, b; j < n; i, j = i+1, j+1 {
		for s[i] == 'a' {
			i++
		}
		if s[i] != s[j] {
			isPossible = false
			break
		}
	}

	bs := n - 1 - ai
	zs := n - a
	if bs < n-a-bs || zs%2 != 0 {
		isPossible = false
	}

	if isPossible {
		fmt.Println(s[:b])
	} else {
		fmt.Println(":(")
	}
}
