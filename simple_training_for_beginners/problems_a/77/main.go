// https://codeforces.com/problemset/problem/432/A

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n, k int
	fmt.Fscan(in, &n, &k)

	validStudents := 0
	for i, x := 0, 0; i < n; i++ {
		fmt.Fscan(in, &x)
		if 5-x >= k {
			validStudents++
		}
	}

	teams := validStudents / 3
	fmt.Println(teams)
}
