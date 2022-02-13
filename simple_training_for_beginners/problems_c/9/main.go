// https://codeforces.com/problemset/problem/1015/C

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type pair struct {
	a, b int64
}

func main() {
	in := bufio.NewReader(os.Stdin)
	var n, m int64
	fmt.Fscan(in, &n, &m)

	ar := make([]pair, n)
	var sum int64
	for i := range ar {
		fmt.Fscan(in, &ar[i].a, &ar[i].b)
		sum += ar[i].b
	}

	if sum > m {
		fmt.Println(-1)
		return
	}

	sort.Slice(ar, func(i, j int) bool {
		return abs(ar[i].a-ar[i].b) < abs(ar[j].a-ar[j].b)
	})

	res := n
	for i := range ar {
		sum += (ar[i].a - ar[i].b)
		if sum <= m {
			res--
		} else {
			break
		}
	}

	fmt.Println(res)
}

func abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}
