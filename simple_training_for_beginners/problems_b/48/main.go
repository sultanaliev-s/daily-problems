// https://codeforces.com/problemset/problem/1197/B

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

	poles := make([]int, n)
	max := 0
	for i := range poles {
		fmt.Fscan(in, &poles[i])
		if poles[max] < poles[i] {
			max = i
		}
	}

	isPossible := true
	for i := max + 1; i < n; i++ {
		if poles[i] > poles[i-1] {
			isPossible = false
		}
	}
	for i := max - 1; i >= 0; i-- {
		if poles[i] > poles[i+1] {
			isPossible = false
		}
	}

	if isPossible {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
