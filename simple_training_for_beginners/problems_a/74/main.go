// https://codeforces.com/problemset/problem/510/A

package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	for i, j := 0, 0; i < n; i++ {
		if i%2 == 0 {
			print('#', m)
		} else {
			if j%2 == 0 {
				print('.', m-1)
				fmt.Print("#")
			} else {
				fmt.Print("#")
				print('.', m-1)
			}
			j++
		}
		fmt.Println()
	}
}

func print(c byte, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c", c)
	}
}
