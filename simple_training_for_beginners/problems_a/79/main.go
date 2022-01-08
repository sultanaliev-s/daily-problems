// https://codeforces.com/problemset/problem/427/A

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(in, &n)

	cops := 0
	untreatedCrimes := 0
	for i, x := 0, 0; i < n; i++ {
		fmt.Fscan(in, &x)
		if x == -1 && cops > 0 {
			cops--
		} else if x == -1 {
			untreatedCrimes++
		} else {
			cops += x
		}
	}

	fmt.Println(untreatedCrimes)
}
