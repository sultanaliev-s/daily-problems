// https://codeforces.com/problemset/problem/41/A

package main

import "fmt"

func main() {
	var berl, birl string
	fmt.Scanln(&berl)
	fmt.Scanln(&birl)

	isCorrect := true
	for i, j := 0, (len(birl) - 1); i < len(berl) && j >= 0; i, j = i+1, j-1 {
		if berl[i] != birl[j] {
			isCorrect = false
			break
		}
	}

	if isCorrect {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
