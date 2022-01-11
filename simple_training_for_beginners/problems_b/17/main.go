// https://codeforces.com/problemset/problem/1248/B

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(in, &n)

	sticks := make([]int, n)
	for i := range sticks {
		fmt.Fscan(in, &sticks[i])
	}
	sort.Ints(sticks)

	vertical := int64(0)
	horizontal := int64(0)
	i := 0
	for ; i < n/2; i++ {
		vertical += int64(sticks[i])
	}
	for ; i < n; i++ {
		horizontal += int64(sticks[i])
	}

	dist := vertical*vertical + horizontal*horizontal

	fmt.Println(dist)
}
