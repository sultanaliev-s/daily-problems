// https://codeforces.com/problemset/problem/110/A

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var number string
	fmt.Scanln(&number)

	var nLuckyNumbers int
	for i := 0; i < len(number); i++ {
		if number[i] == '4' || number[i] == '7' {
			nLuckyNumbers++
		}
	}

	nearlyLuckyNumber := strconv.Itoa(nLuckyNumbers)
	isNearlyLucky := true
	for i := 0; i < len(nearlyLuckyNumber); i++ {
		if nearlyLuckyNumber[i] != '4' && nearlyLuckyNumber[i] != '7' {
			isNearlyLucky = false
		}
	}

	if isNearlyLucky {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
