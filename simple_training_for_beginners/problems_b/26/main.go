// https://codeforces.com/problemset/problem/1080/B

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
		var l, r int
		fmt.Fscan(in, &l, &r)

		length := r - l + 1
		pairs := length / 2
		res := 0
		if length%2 == 0 {
			res = pairs
			if l%2 == 0 {
				res = -res
			}

		} else {
			res = r - pairs
			if l%2 != 0 {
				res = -res
			}
		}

		fmt.Println(res)
	}
}
