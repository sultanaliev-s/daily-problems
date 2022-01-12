// https://codeforces.com/problemset/problem/1095/B

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(in, &n)

	arr := make([]int, n)
	for i := range arr {
		fmt.Fscan(in, &arr[i])
	}

	sort.Ints(arr)

	if arr[n-2]-arr[0] < arr[n-1]-arr[1] {
		fmt.Println(arr[n-2] - arr[0])
	} else {
		fmt.Println(arr[n-1] - arr[1])
	}
}
