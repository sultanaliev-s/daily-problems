// https://codeforces.com/problemset/problem/688/A

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n, d int
	fmt.Fscan(in, &n, &d)

	var enemies string
	var maxWins, wins int
	for i := 0; i < d; i++ {
		fmt.Fscan(in, &enemies)
		if strings.ContainsRune(enemies, rune('0')) {
			wins++
		} else {
			if maxWins < wins {
				maxWins = wins
			}
			wins = 0
		}
	}
	if maxWins < wins {
		maxWins = wins
	}

	fmt.Println(maxWins)
}
