// https://codeforces.com/problemset/problem/1398/B

package main

import (
	"fmt"
	"sort"
)

func main() {
	var t int
	fmt.Scan(&t)

	var line string
	ones := make([]int, 50)
	var count, ind, alice int
	for ; t > 0; t-- {
		fmt.Scan(&line)
		for i := 0; i < len(line); i++ {
			if line[i] == '1' {
				count++
			} else if count > 0 {
				ones[ind] = count
				ind++
				count = 0
			}
		}
		if count > 0 {
			ones[ind] = count
			ind++
		}

		sort.Slice(ones, func(i, j int) bool {
			return ones[i] > ones[j]
		})

		for i := 0; i < ind; i += 2 {
			alice += ones[i]
		}

		fmt.Println(alice)
		count, ind, alice = 0, 0, 0
		for i := range ones {
			ones[i] = 0
		}
	}

}
