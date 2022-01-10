// https://codeforces.com/problemset/problem/1373/B

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

	for ; t > 0; t-- {
		var line string
		fmt.Fscan(in, &line)

		var zeroes, ones int
		for i := range line {
			if line[i] == '0' {
				zeroes++
			} else {
				ones++
			}
		}

		min := ones
		if zeroes < ones {
			min = zeroes
		}

		if min%2 != 0 {
			fmt.Println("DA")
		} else {
			fmt.Println("NET")
		}
	}
}
