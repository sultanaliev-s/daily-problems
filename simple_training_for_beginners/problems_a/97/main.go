// https://codeforces.com/problemset/problem/1343/B

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var t int
	fmt.Fscan(in, &t)

	var n int
	var even, odd int
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &n)
		if n%4 != 0 {
			fmt.Println("NO")
		} else {
			fmt.Println("YES")
			for i, j := 0, 2; i < n/2; i, j = i+1, j+2 {
				fmt.Print(j, " ")
				even += j
			}
			for i, j := 1, 1; i < n/2; i, j = i+1, j+2 {
				fmt.Print(j, " ")
				odd += j
			}
			fmt.Println(even - odd)
			even, odd = 0, 0
		}
	}
}
