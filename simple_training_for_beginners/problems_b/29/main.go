// https://codeforces.com/problemset/problem/999/B

package main

import "fmt"

func main() {
	var n int
	var s string
	fmt.Scan(&n, &s)

	arr := []byte(s)
	for i := 2; i <= n; i++ {
		if n%i == 0 {
			Reverse(arr, i)
		}
	}

	s = string(arr)
	fmt.Println(s)
}

func Reverse(arr []byte, n int) {
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
