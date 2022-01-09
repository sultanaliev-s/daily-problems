// https://codeforces.com/problemset/problem/832/A

package main

import "fmt"

func main() {
	var n, k int64
	fmt.Scan(&n, &k)

	turns := n / k
	if (turns % 2) == 0 {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
	}
}
