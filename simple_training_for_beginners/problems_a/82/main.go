// https://codeforces.com/problemset/problem/266/A

package main

import "fmt"

func main() {
	var (
		n    int
		line string
	)
	fmt.Scan(&n, &line)

	prev := line[0]
	cur := line[0]
	toRemove := 0
	for i := 1; i < n; i++ {
		cur = line[i]
		if cur == prev {
			toRemove++
		}
		prev = cur
	}

	fmt.Println(toRemove)
}
