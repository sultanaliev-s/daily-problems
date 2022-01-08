// https://codeforces.com/problemset/problem/791/A

package main

import "fmt"

func main() {
	var limak, bob int
	fmt.Scan(&limak, &bob)

	years := 0
	for limak <= bob {
		limak *= 3
		bob *= 2
		years++
	}

	fmt.Println(years)
}
