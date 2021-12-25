// https://codeforces.com/problemset/problem/34/A

package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d\n", &n)

	var x, y int
	fmt.Scanf("%d %d", &x, &y)
	first := x
	xi, yi := 0, 1
	minDiff := Abs(x - y)
	x = y
	for i := 2; i < n; i++ {
		fmt.Scanf("%d", &y)
		if Abs(x-y) < minDiff {
			minDiff = Abs(x - y)
			xi, yi = i-1, i
		}
		x = y
	}
	if Abs(first-y) < minDiff {
		xi, yi = n-1, 0
	}

	fmt.Println(xi+1, yi+1)

}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
