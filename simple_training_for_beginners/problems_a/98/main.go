// https://codeforces.com/problemset/problem/1154/A

package main

import (
	"fmt"
	"sort"
)

func main() {
	var nums [4]int
	fmt.Scan(&nums[0], &nums[1], &nums[2], &nums[3])
	sort.Ints(nums[:])

	a := nums[3] - nums[0]
	b := nums[3] - nums[1]
	c := nums[3] - nums[2]
	fmt.Println(a, b, c)
}
