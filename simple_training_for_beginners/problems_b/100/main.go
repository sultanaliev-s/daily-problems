// https://codeforces.com/problemset/problem/362/B

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n, m int
	fmt.Fscan(in, &n, &m)

	dirty := make([]int, m)
	for i := range dirty {
		fmt.Fscan(in, &dirty[i])
	}

	sort.Ints(dirty)
	isPossible := true
	if m > 0 && (dirty[0] == 1 || dirty[m-1] == n) {
		isPossible = false
	} else {
		streak := 0
		for i := 0; i < m-1; i++ {
			if dirty[i]+1 == dirty[i+1] {
				streak++
			} else {
				streak = 0
			}
			if streak == 2 {
				isPossible = false
				break
			}
		}
	}

	if isPossible {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
