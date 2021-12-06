// https://codeforces.com/problemset/problem/177/A2

package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d\n", &n)

	var matrix = make([][]int, n)
	for i := range matrix {
		var arr = make([]int, n)
		for j := range arr {
			fmt.Scanf("%d", &arr[j])
		}
		fmt.Scanln()
		matrix[i] = arr
	}

	var sum int
	var middle = n / 2
	for i, j := 0, n-1; i < n; i, j = i+1, j-1 {
		sum += matrix[i][middle]
		sum += matrix[middle][i]
		sum += matrix[i][i]
		sum += matrix[i][j]
	}
	sum -= matrix[middle][middle] * 3

	fmt.Println(sum)
}
