// https://codeforces.com/problemset/problem/80/A

package main

import (
	"fmt"
	"math"
)

func isPrime(x int) bool {
	if x == 2 {
		return true
	}
	sqrtX := int(math.Sqrt(float64(x)))
	for i := 2; i <= sqrtX; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var n, m int
	fmt.Scanf("%d %d\n", &n, &m)

	var i int
	for i = n + 1; i <= m; i++ {
		if isPrime(i) {
			break
		}
	}
	if i == m {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
