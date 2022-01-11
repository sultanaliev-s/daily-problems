// https://codeforces.com/problemset/problem/1337/B

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
		var x, n, m int
		fmt.Fscan(in, &x, &n, &m)

		hpAfterAbsorption := castAbsorptionsOptimally(x, n)
		resHp := hpAfterAbsorption - m*10

		if resHp > 0 {
			fmt.Println("NO")
		} else {
			fmt.Println("YES")
		}
	}
}

func castAbsorptionsOptimally(x, n int) int {
	var t int
	for i := 0; i < n; i++ {
		t = x/2 + 10
		if t < x {
			x = t
		}
	}
	return x
}
