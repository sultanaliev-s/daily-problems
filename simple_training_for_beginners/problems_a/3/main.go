// https://codeforces.com/problemset/problem/236/A

package main

import "fmt"

func main() {
	var nickname string
	fmt.Scanf("%s\n", &nickname)

	var symbols = make(map[rune]struct{})
	for _, c := range nickname {
		symbols[c] = struct{}{}
	}

	if isEven := (len(symbols) % 2) == 0; isEven {
		fmt.Println("CHAT WITH HER!")
	} else {
		fmt.Println("IGNORE HIM!")
	}
}
