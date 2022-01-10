// https://codeforces.com/problemset/problem/1374/B

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
		var n int64
		fmt.Fscan(in, &n)
		var cnt2, cnt3 int
		for n%2 == 0 {
			n /= 2
			cnt2++
		}
		for n%3 == 0 {
			n /= 3
			cnt3++
		}

		if n == 1 && cnt3 >= cnt2 {
			fmt.Println(cnt3*2 - cnt2)
		} else {
			fmt.Println(-1)
		}
	}
}
