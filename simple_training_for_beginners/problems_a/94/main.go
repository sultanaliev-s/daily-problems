// https://codeforces.com/problemset/problem/1400/A

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

	var n int
	var line string
	for test := 0; test < t; test++ {
		fmt.Fscan(in, &n, &line)
		for i := 0; i < len(line); i += 2 {
			fmt.Printf("%c", line[i])
		}
		fmt.Println()
	}
}
