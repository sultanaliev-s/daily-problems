// https://codeforces.com/problemset/problem/1099/C

package main

import "fmt"

func main() {
	var s string
	var k int
	fmt.Scan(&s, &k)

	snow := 0
	pops := 0

	for i := range s {
		switch s[i] {
		case '*':
			snow++
		case '?':
			pops++
		}
	}

	letters := len(s) - snow - pops
	if letters-k > pops+snow || (k-letters > 0 && snow == 0) {
		fmt.Println("Impossible")
	} else {
		ar := make([]byte, 0, len(s))
		ind := 0
		toDel := 0
		toProlong := 0
		if letters-k > 0 {
			toDel = letters - k
		} else if k-letters > 0 {
			toProlong = k - letters
		}
		for i := range s {
			switch s[i] {
			case '*':
				for toProlong > 0 {
					ar = append(ar, s[i-1])
					ind++
					toProlong--
				}
				if toDel > 0 {
					toDel--
					ind--
					ar = ar[:ind]
				}
			case '?':
				if toDel > 0 {
					toDel--
					ind--
					ar = ar[:ind]
				}
			default:
				ar = append(ar, s[i])
				ind++
			}
		}
		for i := range ar {
			fmt.Printf("%c", ar[i])
		}
		fmt.Println()
	}
}
