// https://codeforces.com/problemset/problem/282/A

package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	x := 0
	var cmd string
	for i := 0; i < n; i++ {
		fmt.Scan(&cmd)
		for j := range cmd {
			if cmd[j] == '+' {
				x++
				break
			} else if cmd[j] == '-' {
				x--
				break
			}
		}
	}

	fmt.Println(x)
}
