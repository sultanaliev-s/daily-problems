// https://codeforces.com/problemset/problem/125/B

package main

import "fmt"

func Indent(x int) {
	for i := 0; i < x; i++ {
		fmt.Print("  ")
	}
}

func main() {
	var s string
	fmt.Scan(&s)

	ops, ends := 0, 0
	b, e := 0, 0
	for i := range s {
		if s[i] == '<' {
			b = i
		} else if s[i] == '>' {
			e = i
			if s[b+1] == '/' {
				ends++
				Indent(ops - ends)
			} else {
				Indent(ops - ends)
				ops++
			}
			fmt.Println(s[b : e+1])
		}
	}
}
