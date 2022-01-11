// https://codeforces.com/problemset/problem/1206/B

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

	var res int64
	var x, pos, neg, zeroes int64
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &x)
		if x == 0 {
			zeroes++
		} else if x < 0 {
			neg++
			res += -1 - x
		} else {
			pos++
			res += x - 1
		}
	}

	if neg%2 != 0 && zeroes == 0 {
		res += 2
	}
	res += zeroes

	fmt.Println(res)
}
