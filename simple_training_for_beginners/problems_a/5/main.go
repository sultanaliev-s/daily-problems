// https://codeforces.com/problemset/problem/214/A

package main

import (
	"fmt"
	"math"
)

func main() {
	var n, m int
	fmt.Scanf("%d %d\n", &n, &m)

	var nSolutions int

	var sqrtN int = int(math.Sqrt(float64(n))) + 1

	for i := 0; i <= sqrtN; i++ {
		for j := 0; j <= n; j++ {
			if (i*i+j) == n && (i+j*j) == m {
				nSolutions++
			}
		}
	}

	fmt.Println(nSolutions)
}
