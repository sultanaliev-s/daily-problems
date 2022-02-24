// https://codeforces.com/problemset/problem/1214/C

package main

import "fmt"

func main() {
	var n int
	var s string
	fmt.Scan(&n, &s)

	open := 0
	close := 0
	hasMoved := false
	isPossible := true
	for i := range s {
		if s[i] == '(' {
			open++
		} else {
			close++
			if close > open && !hasMoved {
				hasMoved = true
			} else if close-1 > open && hasMoved {
				isPossible = false
				break
			}
		}
	}

	if open != close {
		isPossible = false
	}

	if isPossible {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
