// https://codeforces.com/problemset/problem/61/A

package main

import "fmt"

func main() {
	var a, b string
	fmt.Scanln(&a)
	fmt.Scanln(&b)

	c := make([]byte, len(a))
	for i := range a {
		if a[i] != b[i] {
			c[i] = '1'
		} else {
			c[i] = '0'
		}
	}

	fmt.Println(string(c))
}
