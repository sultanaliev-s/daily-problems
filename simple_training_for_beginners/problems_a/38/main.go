// https://codeforces.com/problemset/problem/4/A

package main

import "fmt"

func main() {
	var w int
	fmt.Scanf("%d\n", &w)

	if w%2 == 0 && w != 2 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
