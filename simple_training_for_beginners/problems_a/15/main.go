// https://codeforces.com/problemset/problem/146/A

package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanf("%d\n", &n)
	var ticket string
	fmt.Scanln(&ticket)

	var arr = []byte(ticket)
	for i, v := range arr {
		arr[i] = v - '0'
	}

	var isLucky = false

	if containsOnlyFourAndSeven(&arr) {
		firstHalf := sum(&arr, 0, n/2)
		secondHalf := sum(&arr, n/2, n)

		isLucky = firstHalf == secondHalf
	}

	if isLucky {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func containsOnlyFourAndSeven(arr *[]byte) bool {
	for _, v := range *arr {
		if v != 4 && v != 7 {
			return false
		}
	}
	return true
}

func sum(arr *[]byte, start, end int) (res int) {
	for i := start; i < end; i++ {
		res += int((*arr)[i])
	}

	return res
}
