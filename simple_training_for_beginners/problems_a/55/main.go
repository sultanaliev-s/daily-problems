// https://codeforces.com/problemset/problem/703/A

package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	mishka := 0
	chris := 0
	var m, c int
	for i := 0; i < n; i++ {
		fmt.Scan(&m, &c)
		if m > c {
			mishka++
		} else if m < c {
			chris++
		}
	}

	if mishka > chris {
		fmt.Println("Mishka")
	} else if mishka < chris {
		fmt.Println("Chris")
	} else {
		fmt.Println("Friendship is magic!^^")
	}
}
