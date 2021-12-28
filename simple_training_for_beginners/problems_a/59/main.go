// https://codeforces.com/problemset/problem/686/A

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n int
	var iceCreams uint64
	fmt.Fscan(in, &n, &iceCreams)

	sadKids := 0
	op := ""
	x := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &op, &x)
		switch op {
		case "+":
			iceCreams += uint64(x)
		case "-":
			if uint64(x) > iceCreams {
				sadKids++
			} else {
				iceCreams -= uint64(x)
			}
		}
	}

	fmt.Println(iceCreams, sadKids)
}
