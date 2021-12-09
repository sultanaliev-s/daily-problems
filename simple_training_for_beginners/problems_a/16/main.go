// https://codeforces.com/problemset/problem/144/A

package main

import "fmt"

func main() {
	var n, x int
	fmt.Scanf("%d\n", &n)

	fmt.Scanf("%d", &x)
	min, max := x, x
	minInd, maxInd := 0, 0

	for i := 1; i < n; i++ {
		fmt.Scanf("%d", &x)
		if x > max {
			max = x
			maxInd = i
		}
		if x <= min {
			min = x
			minInd = i
		}
	}

	var res int
	if maxInd > minInd {
		res -= 1
	}

	res += maxInd + (n - minInd - 1)

	fmt.Println(res)
}
