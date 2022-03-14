// https://codeforces.com/problemset/problem/1326/C

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

	type pair struct {
		val int
		pos int
	}
	ar := make([]pair, n)
	for i := range ar {
		fmt.Fscan(in, &ar[i].val)
		ar[i].pos = i
	}

	sort.Slice(ar, func(i, j int) bool { return ar[i].val > ar[j].val })

	poss := make([]int, k)
	sum := int64(0)
	for i := 0; i < k; i++ {
		poss[i] = ar[i].pos
		sum += int64(ar[i].val)
	}

	sort.Ints(poss)

	dist := make([]int64, k)
	for i := 0; i < k-1; i++ {
		dist[i] = int64(poss[i+1] - poss[i])
	}
	dist[k-1] = 1

	res := dist[0]
	for i := 1; i < k; i++ {
		res = res * dist[i] % 998244353
	}

	fmt.Println(sum, res)
}
