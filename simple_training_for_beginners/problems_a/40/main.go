// https://codeforces.com/problemset/problem/758/A

package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d\n", &n)

	citizensMoney := make([]int, n)
	max := 0
	for i := range citizensMoney {
		fmt.Scanf("%d", &citizensMoney[i])
		if citizensMoney[i] > max {
			max = citizensMoney[i]
		}
	}

	moneyNeededToBalance := 0
	for _, v := range citizensMoney {
		moneyNeededToBalance += max - v
	}

	fmt.Println(moneyNeededToBalance)
}
