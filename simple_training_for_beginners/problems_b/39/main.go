// https://codeforces.com/problemset/problem/1312/B

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	var t int
	fmt.Fscan(in, &t)

	for ; t > 0; t-- {
		var n int
		fmt.Fscan(in, &n)

		nums := make([]int, n)
		for i := range nums {
			fmt.Fscan(in, &nums[i])
		}

		sort.Slice(nums, func(i, j int) bool {
			return nums[i] > nums[j]
		})

		for i := range nums {
			if i != 0 {
				fmt.Fprint(out, " ")
			}
			fmt.Fprint(out, nums[i])
		}

		fmt.Fprintln(out)
		out.Flush()
	}
}
