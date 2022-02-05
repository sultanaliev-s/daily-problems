// https://codeforces.com/problemset/problem/137/B

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

	ar := make([]int, 5001)
	var x int
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &x)
		ar[x]++
	}

	res := 0
	for i := 1; i < n+1; i++ {
		if ar[i] == 0 {
			res++
		}
	}

	fmt.Println(res)
}
