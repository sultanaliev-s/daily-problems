// https://codeforces.com/problemset/problem/228/A

package main

import (
	"fmt"
	"sort"
)

func main() {
	var horseshoes [4]int
	fmt.Scanf("%d %d %d %d\n",
		&horseshoes[0], &horseshoes[1], &horseshoes[2], &horseshoes[3])

	sort.Ints(horseshoes[:])
	var nSame int
	for i := 1; i < 4; i++ {
		if horseshoes[i-1] == horseshoes[i] {
			nSame++
		}
	}

	fmt.Println(nSame)

}
