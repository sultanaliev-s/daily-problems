// https://codeforces.com/problemset/problem/116/A

package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d\n", &n)

	var a, b, capacity, minCapacity int
	for i := 0; i < n; i++ {
		fmt.Scanf("%d %d\n", &a, &b)
		capacity += b - a
		if capacity > minCapacity {
			minCapacity = capacity
		}
	}

	fmt.Println(minCapacity)
}
