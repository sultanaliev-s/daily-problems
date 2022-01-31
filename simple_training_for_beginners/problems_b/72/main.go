// https://codeforces.com/problemset/problem/652/B

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var n int
	fmt.Fscan(in, &n)

	ar := make([]int, n)
	for i := range ar {
		fmt.Fscan(in, &ar[i])
	}

	sort.Ints(ar)

	res := make([]int, n)
	var i, j, k int
	for i, j, k = 0, 0, n-1; i < n; i++ {
		if i%2 == 0 {
			res[i] = ar[j]
			j++
		} else {
			res[i] = ar[k]
			k--
		}
	}

	for i := range res {
		if i > 0 {
			fmt.Fprint(out, " ")
		}
		fmt.Fprint(out, res[i])
	}

}
