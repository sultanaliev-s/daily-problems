// https://codeforces.com/problemset/problem/981/B

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	nums := make(map[int]int)

	var n, x, y int
	var sum int64
	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &x, &y)
		sum += int64(y)
		nums[x] = y
	}
	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &x, &y)
		if t := nums[x]; t < y {
			sum += int64(y - t)
		}
	}

	fmt.Println(sum)
}
