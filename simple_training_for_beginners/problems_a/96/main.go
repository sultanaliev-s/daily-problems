// https://codeforces.com/problemset/problem/1369/A

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
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &n)
		if n%4 == 0 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
