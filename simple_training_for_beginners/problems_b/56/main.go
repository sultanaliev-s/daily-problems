// https://codeforces.com/problemset/problem/1073/B

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	var n int
	fmt.Fscan(in, &n)

	stack := make([]int, n)
	order := make([]int, n+1)
	for i := range stack {
		fmt.Fscan(in, &stack[i])
		order[stack[i]] = i
	}

	var x, ind int
	isFirst := true
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &x)

		res := 0
		if order[x] >= ind {
			res += order[x] + 1 - ind
			ind = order[x] + 1
		}

		if !isFirst {
			fmt.Fprint(out, " ")
		}
		isFirst = false
		fmt.Fprint(out, res)
	}
	fmt.Fprintln(out)
	out.Flush()
}
