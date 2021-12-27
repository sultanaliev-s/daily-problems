// https://codeforces.com/problemset/problem/734/A

package main

import "fmt"

func main() {
	var n int
	var line string
	fmt.Scanf("%d\n%s\n", &n, &line)

	anton := 0
	danik := 0
	for _, v := range line {
		if v == 'A' {
			anton++
		} else {
			danik++
		}
	}

	if anton == danik {
		fmt.Println("Friendship")
	} else if anton > danik {
		fmt.Println("Anton")
	} else {
		fmt.Println("Danik")
	}
}
