// https://codeforces.com/problemset/problem/672/A

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)

	num := 1
	i := 1
	for {
		if i > n {
			num--
			break
		}

		if num < 10 {
			i++
		} else if num < 100 {
			i += 2
		} else if num < 1000 {
			i += 3
		} else {
			i += 4
		}
		num++
	}

	nums := strconv.Itoa(num)
	fmt.Printf("%c\n", nums[len(nums)-(i-n)])
}
