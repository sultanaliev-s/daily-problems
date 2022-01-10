// https://codeforces.com/problemset/problem/1367/B

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

	for ; t > 0; t-- {
		var n int
		fmt.Fscan(in, &n)

		neededEven := 0
		neededOdd := 0
		var x int
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &x)
			if i%2 == 0 && x%2 != 0 {
				neededEven++
			} else if i%2 != 0 && x%2 == 0 {
				neededOdd++
			}
		}

		if neededEven == neededOdd {
			fmt.Println(neededEven)
		} else {
			fmt.Println(-1)
		}
	}
}
