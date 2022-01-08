// https://codeforces.com/problemset/problem/785/A

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

	faces := 0
	figure := ""
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &figure)
		switch figure[0] {
		case 'T':
			faces += 4
		case 'C':
			faces += 6
		case 'O':
			faces += 8
		case 'D':
			faces += 12
		case 'I':
			faces += 20
		}
	}

	fmt.Println(faces)
}
