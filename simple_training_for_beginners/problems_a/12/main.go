// https://codeforces.com/problemset/problem/1399/A

package main

import (
	"fmt"
	"sort"
)

func main() {
	var nTests int
	fmt.Scanf("%d\n", &nTests)

	for test := 0; test < nTests; test++ {
		var arrLength int
		fmt.Scanf("%d\n", &arrLength)

		var arr = make([]int, arrLength)

		for i := 0; i < arrLength; i++ {
			fmt.Scanf("%d", &arr[i])
		}

		sort.Ints(arr)
		var impossibles int
		for i := 1; i < arrLength; i++ {
			if (arr[i]-arr[i-1]) < -1 || (arr[i]-arr[i-1]) > 1 {
				impossibles++
			}
		}
		fmt.Scanln()

		if impossibles > 0 {
			fmt.Println("NO")
		} else {
			fmt.Println("YES")
		}
	}
}
