// https://codeforces.com/problemset/problem/59/A

package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	var line string
	fmt.Scanln(&line)

	var lowers, uppers int
	for _, v := range line {
		if unicode.IsUpper(v) {
			uppers++
		} else {
			lowers++
		}
	}

	if uppers > lowers {
		fmt.Println(strings.ToUpper(line))
	} else {
		fmt.Println(strings.ToLower(line))
	}
}
