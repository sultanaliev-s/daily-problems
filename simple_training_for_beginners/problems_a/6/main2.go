// https://codeforces.com/problemset/problem/202/A

package main

import (
	"fmt"
	"sort"
)

func main() {
	var line string
	fmt.Scanln(&line)

	symbols := []byte(line)
	sort.Slice(symbols, func(i int, j int) bool {
		return symbols[i] > symbols[j]
	})

	for _, v := range symbols {
		if v == symbols[0] {
			fmt.Printf("%c", v)
		}
	}

	fmt.Println()
}
