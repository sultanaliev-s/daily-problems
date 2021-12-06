// https://codeforces.com/problemset/problem/202/A

package main

import (
	"fmt"
	"sort"
)

func main() {
	var line string
	fmt.Scanln(&line)

	var combinations []string
	for comboLength := 0; comboLength <= len(line); comboLength++ {
		var used = make([]bool, len(line))
		var str string
		makeCombinations(0, comboLength, &used, str, &combinations, &line)
	}

	sort.Strings(combinations)
	fmt.Println(combinations[len(combinations)-1])
}

func makeCombinations(currentIndex int, comboLength int, used *[]bool,
	str string, combinations *[]string, line *string) {
	if len(str) >= comboLength {
		if isPalindrome(str) {
			*combinations = append(*combinations, str)
		}
		return
	}

	for i := currentIndex; i < len(*line); i++ {
		if !(*used)[i] {
			(*used)[i] = true
			makeCombinations(i+1, comboLength, used,
				str+string((*line)[i]), combinations, line)
			(*used)[i] = false
		}
	}

}

func isPalindrome(str string) bool {
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		if str[i] != str[j] {
			return false
		}
	}
	return true
}
