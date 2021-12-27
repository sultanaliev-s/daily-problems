// https://codeforces.com/problemset/problem/707/A

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n, m int
	fmt.Fscan(in, &n, &m)

	isBlackWhite := true
	char := ""
	for i := 0; i < n*m; i++ {
		fmt.Fscan(in, &char)
		if char != "B" && char != "W" && char != "G" {
			isBlackWhite = false
		}
	}

	if isBlackWhite {
		fmt.Println("#Black&White")
	} else {
		fmt.Println("#Color")
	}
}
