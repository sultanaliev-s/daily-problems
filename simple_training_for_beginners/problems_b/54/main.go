// https://codeforces.com/problemset/problem/1088/B

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n, k int
	fmt.Fscan(in, &n, &k)

	ar := make([]int, n)
	for i := range ar {
		fmt.Fscan(in, &ar[i])
	}

	sort.Ints(ar)

	sub := 0
	for i, j := 0, 0; i < k; {
		if j < n && ar[j]-sub > 0 {
			fmt.Println(ar[j] - sub)
			sub += ar[j] - sub
			i++
		} else if j < n {
			j++
		} else {
			fmt.Println(0)
			i++
		}
	}
}
