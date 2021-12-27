// https://codeforces.com/problemset/problem/705/A

package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	hate := "I hate "
	love := "I love "
	that := "that "
	for i := 1; i <= n; i++ {
		if i > 1 {
			fmt.Print(that)
		}

		if (i % 2) == 0 {
			fmt.Print(love)
		} else {
			fmt.Print(hate)
		}
	}

	fmt.Println("it")
}
