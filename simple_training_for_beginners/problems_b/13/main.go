// https://codeforces.com/problemset/problem/1335/B

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

	alphabet := "abcdefghijklmnopqrstuvwxyz"
	for ; t > 0; t-- {
		var n, a, b int
		fmt.Fscan(in, &n, &a, &b)
		for i := 0; i < n/b; i++ {
			fmt.Print(alphabet[:b])
		}
		if n%b != 0 {
			fmt.Println(alphabet[:n%b])
		} else {
			fmt.Println()
		}
	}
}
