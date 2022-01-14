// https://codeforces.com/problemset/problem/991/B

package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	grades := make([]int, n)
	sum := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&grades[i])
		sum += grades[i]
	}

	sort.Ints(grades)

	labs := 0
	for i := 0; i < n; i++ {
		if isEnough(sum, n) {
			break
		}
		sum += 5 - grades[i]
		labs++
	}

	fmt.Println(labs)
}

func isEnough(sum, n int) bool {
	avg := float64(sum) / float64(n)
	grade := int(math.Floor(avg + 0.5))
	if grade == 5 {
		return true
	}
	return false
}
