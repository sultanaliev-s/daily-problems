// https://codeforces.com/problemset/problem/681/A

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(in, &n)

	var handle string
	var before, after int
	hasPerformedWell := false
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &handle, &before, &after)
		if 2400 <= before && before < after {
			hasPerformedWell = true
		}
	}

	if hasPerformedWell {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
