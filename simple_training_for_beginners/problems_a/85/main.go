// https://codeforces.com/problemset/problem/339/A

package main

import (
	"fmt"
	"sort"
)

func main() {
	var line string
	fmt.Scan(&line)

	nums := make([]int, len(line)/2+1)
	for i, j := 0, 0; i < len(line); i += 2 {
		nums[j] = int(line[i] - '0')
		j++
	}

	sort.Ints(nums)

	fmt.Print(nums[0])
	for i := 1; i < len(nums); i++ {
		fmt.Print("+", nums[i])
	}
	fmt.Println()
}
