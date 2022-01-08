// https://codeforces.com/problemset/problem/281/A

package main

import "fmt"

func main() {
	var word string
	fmt.Scan(&word)
	word = capitalize(word)
	fmt.Println(word)
}

func capitalize(word string) string {
	if word == "" || !('a' <= word[0] && word[0] <= 'z') {
		return word
	}
	return string(word[0]-('a'-'A')) + word[1:]
}
