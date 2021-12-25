// https://codeforces.com/problemset/problem/755/A

package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scanf("%d\n", &n)

	for i := 1; i <= 1000; i++ {
		if !IsPrime(n*i + 1) {
			fmt.Println(i)
			break
		}
	}
}

func IsPrime(x int) bool {
	if x == 2 {
		return true
	}

	sqrt := int(math.Sqrt(float64(x)))
	for i := 2; i <= sqrt; i++ {
		if (x % i) == 0 {
			return false
		}
	}

	return true
}
