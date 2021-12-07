// https://codeforces.com/problemset/problem/155/A

package main

import "fmt"

func main() {
	var nContests int
	fmt.Scanf("%d\n", &nContests)

	var x int
	fmt.Scanf("%d", &x)
	var max, min = x, x
	var amazings int
	for i := 1; i < nContests; i++ {
		fmt.Scanf("%d", &x)
		if x > max {
			max = x
			amazings++
		}
		if x < min {
			min = x
			amazings++
		}
	}

	fmt.Println(amazings)
}
