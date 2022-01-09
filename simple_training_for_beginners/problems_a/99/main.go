// https://codeforces.com/problemset/problem/1146/A

package main

import "fmt"

func main() {
	var line string
	fmt.Scan(&line)

	nA := 0
	for i := range line {
		if line[i] == 'a' {
			nA++
		}
	}

	maxLength := len(line)
	if nA <= len(line)/2 {
		maxLength = nA + nA - 1
	}

	fmt.Println(maxLength)
}
