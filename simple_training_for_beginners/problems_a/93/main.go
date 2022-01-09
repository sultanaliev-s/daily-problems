// https://codeforces.com/problemset/problem/908/A

package main

import "fmt"

func main() {
	var line string
	fmt.Scan(&line)

	checks := 0
	for i := range line {
		if isNumber(line[i]) {
			if (line[i]-'0')%2 != 0 {
				checks++
			}
		} else if isVowel(line[i]) {
			checks++
		}
	}

	fmt.Println(checks)
}

func isNumber(c byte) bool {
	return '0' <= c && c <= '9'
}

func isVowel(c byte) bool {
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
}
