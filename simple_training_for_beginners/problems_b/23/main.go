// https://codeforces.com/problemset/problem/1141/B

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(in, &n)

	restHours := make([]int, 0)
	var resting, max, x int
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &x)
		if x == 1 {
			resting++
		} else {
			restHours = append(restHours, resting)
			if max < resting {
				max = resting
			}
			resting = 0
		}
	}
	if resting > 0 {
		restHours[0] += resting
		resting = 0
	}

	if max < restHours[0] {
		max = restHours[0]
	}

	fmt.Println(max)
}
