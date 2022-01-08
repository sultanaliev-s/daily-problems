// https://codeforces.com/problemset/problem/1131/A

package main

import "fmt"

func main() {
	var w1, h1, w2, h2 int
	fmt.Scan(&w1, &h1, &w2, &h2)

	res := (h1+h2+2)*2 + w1*2
	fmt.Println(res)
}
