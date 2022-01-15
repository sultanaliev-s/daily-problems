// https://codeforces.com/problemset/problem/1165/B

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

	contests := make([]int, n)
	for i := range contests {
		fmt.Fscan(in, &contests[i])
	}

	sort.Ints(contests)

	day := 1
	for i := 0; i < n; i++ {
		for i < n && contests[i] < day {
			i++
		}
		if i < n && contests[i] >= day {
			day++
		}
	}
	day--
	fmt.Println(day)
}
