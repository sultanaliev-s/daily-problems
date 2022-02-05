// https://codeforces.com/problemset/problem/160/B

package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	var str string
	fmt.Scan(&n, &str)

	left := make([]int, n)
	right := make([]int, n)
	for i := range left {
		left[i] = int(str[i] - '0')
	}
	for i := range right {
		right[i] = int(str[i+n] - '0')
	}

	sort.Ints(left)
	sort.Ints(right)

	isUnlucky := true
	if left[0] < right[0] {
		for i := range left {
			if left[i] >= right[i] {
				isUnlucky = false
				break
			}
		}
	} else if left[0] > right[0] {
		for i := range right {
			if left[i] <= right[i] {
				isUnlucky = false
				break
			}
		}
	} else {
		isUnlucky = false
	}

	if isUnlucky {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
