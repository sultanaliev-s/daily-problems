// https://codeforces.com/problemset/problem/49/A

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

	vowels := []byte{'a', 'e', 'i', 'o', 'u', 'y'}
	isLastLetterVowel := false
	for i := len(line) - 1; i >= 0; i-- {
		if IsLetter(line[i]) {
			if Contains(&vowels, GetLower(line[i])) {
				isLastLetterVowel = true
			}
			break
		}
	}

	if isLastLetterVowel {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func IsLetter(c byte) bool {
	if ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z') {
		return true
	}

	return false
}

func GetLower(c byte) byte {
	if 'A' <= c && c <= 'Z' {
		return c + 'a' - 'A'
	}

	return c
}

func Contains(arr *[]byte, c byte) bool {
	for _, v := range *arr {
		if v == c {
			return true
		}
	}

	return false
}
