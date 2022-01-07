// https://codeforces.com/problemset/problem/443/A

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	line := scanner.Text()
	set := line[1:(len(line) - 1)]
	var letters [26]bool
	for i := 0; i < len(set); i += 3 {
		letters[set[i]-'a'] = true
	}

	count := 0
	for i := range letters {
		if letters[i] {
			count++
		}
	}

	fmt.Println(count)
}
