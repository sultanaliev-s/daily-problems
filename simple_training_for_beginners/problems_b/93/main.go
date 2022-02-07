// https://codeforces.com/problemset/problem/1370/B

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var t int
	fmt.Fscan(in, &t)
	for ; t > 0; t-- {
		var n int
		fmt.Fscan(in, &n)

		arr := make([]int, n*2)
		for i := range arr {
			fmt.Fscan(in, &arr[i])
		}

		for i, p := 0, 1; i < n*2 && p < n; i++ {
			if arr[i] != 0 {
				for j := 0; j < n*2; j++ {
					if j != i && arr[j] != 0 && (arr[i]+arr[j])%2 == 0 {
						fmt.Println(i+1, j+1)
						arr[i], arr[j] = 0, 0
						p++
						break
					}
				}
			}
		}
	}
}
