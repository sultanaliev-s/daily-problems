// https://codeforces.com/problemset/problem/431/A

package main

import "fmt"

func main() {
	var a1, a2, a3, a4 int
	var game string
	fmt.Scan(&a1, &a2, &a3, &a4, &game)

	sumOfCalories := 0
	for i := range game {
		switch game[i] {
		case '1':
			sumOfCalories += a1
		case '2':
			sumOfCalories += a2
		case '3':
			sumOfCalories += a3
		case '4':
			sumOfCalories += a4
		}
	}

	fmt.Println(sumOfCalories)
}
