// https://codeforces.com/problemset/problem/766/B

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

	nums := make([]int, n)
	for i := range nums {
		fmt.Fscan(in, &nums[i])
	}

	sort.Ints(nums)

	isPossible := false
	for i := 0; i < n-2; i++ {
		a := nums[i] + nums[i+1]
		b := nums[i] + nums[i+2]
		c := nums[i+1] + nums[i+2]

		if a > nums[i+2] && b > nums[i+1] && c > nums[i] {
			isPossible = true
			break
		}
	}

	if isPossible {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
