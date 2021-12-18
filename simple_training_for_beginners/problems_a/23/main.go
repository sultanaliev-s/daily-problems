// https://codeforces.com/problemset/problem/104/A

package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d\n", &n)

	required := n - 10
	if required < 1 || required > 11 {
		fmt.Println(0)
	} else if required == 10 {
		fmt.Println(15)
	} else {
		fmt.Println(4)
	}
}
