// https://codeforces.com/problemset/problem/746/B

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n int
	var line string
	fmt.Fscan(in, &n, &line)

	word := make([]byte, n)
	for i, j := n-1, n-1; j >= 0; i, j = i-1, j-2 {
		word[i] = line[j]
	}
	for i, j := 0, n-2; j >= 0; i, j = i+1, j-2 {
		word[i] = line[j]
	}

	fmt.Println(string(word))
}
