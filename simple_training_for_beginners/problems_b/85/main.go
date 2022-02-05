// https://codeforces.com/problemset/problem/122/B

package main

import "fmt"

func main() {
	var str string
	fmt.Scan(&str)

	f := 0
	s := 0
	for i := range str {
		if str[i] == '4' {
			f++
		} else if str[i] == '7' {
			s++
		}
	}

	if s+f == 0 {
		fmt.Println(-1)
	} else if s > f {
		fmt.Println(7)
	} else {
		fmt.Println(4)
	}
}
