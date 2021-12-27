// https://codeforces.com/problemset/problem/749/A

package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d\n", &n)

	twos := n / 2
	threes := 0

	if (n % 2) != 0 {
		twos -= 1
		threes = 1
	}

	fmt.Println(twos + threes)
	for i := 0; i < twos; i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(2)
	}

	if threes > 0 {
		fmt.Print(" ", 3)
	}

	fmt.Println()
}
