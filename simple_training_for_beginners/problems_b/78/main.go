// https://codeforces.com/problemset/problem/284/B

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n int
	var s string
	fmt.Fscan(in, &n, &s)

	ni := 0
	na := 0
	for i := range s {
		switch s[i] {
		case 'I':
			ni++
		case 'A':
			na++
		}
	}

	if ni > 1 {
		fmt.Println(0)
	} else if ni == 1 {
		fmt.Println(ni)
	} else {
		fmt.Println(na)
	}
}
