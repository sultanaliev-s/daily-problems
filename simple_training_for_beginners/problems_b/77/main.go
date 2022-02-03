// https://codeforces.com/problemset/problem/339/B

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n, m int64
	fmt.Fscan(in, &n, &m)

	var res, id, x int64
	id = 1
	for i := int64(0); i < m; i++ {
		fmt.Fscan(in, &x)
		if id <= x {
			res += x - id
		} else {
			res += n - id + x
		}
		id = x
	}

	fmt.Println(res)
}
