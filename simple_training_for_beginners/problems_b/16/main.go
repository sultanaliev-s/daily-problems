// https://codeforces.com/problemset/problem/1283/B

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
		var sweets, children, toGive int
		fmt.Fscan(in, &sweets, &children)

		toGive = sweets
		sweets = sweets % children
		toGive -= sweets - Min(sweets, children/2)
		fmt.Println(toGive)
	}
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
