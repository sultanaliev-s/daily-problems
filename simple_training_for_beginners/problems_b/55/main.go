// https://codeforces.com/problemset/problem/1077/B

package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	flats := make([]int, n)
	for i := range flats {
		fmt.Scan(&flats[i])
	}

	count := 0
	for i := 1; i < n-1; i++ {
		if flats[i] == 0 && flats[i-1]+flats[i+1] == 2 {
			count++
			flats[i+1] = 0
		}
	}

	fmt.Println(count)
}
