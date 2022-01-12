// https://codeforces.com/problemset/problem/1093/B

package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)

	for ; t > 0; t-- {
		var line string
		fmt.Scan(&line)

		if !isPalindrome(line) {
			fmt.Println(line)
		} else {
			i := 1
			for ; i < len(line); i++ {
				if line[i-1] != line[i] {
					break
				}
			}
			if i != len(line) {
				fmt.Print(line[:i-1], line[i:i+1], line[i-1:i], line[i+1:])
				fmt.Println()
			} else {
				fmt.Println(-1)
			}
		}
	}
}

func isPalindrome(line string) bool {
	for i, j := 0, len(line)-1; i < j; i, j = i+1, j-1 {
		if line[i] != line[j] {
			return false
		}
	}
	return true
}
