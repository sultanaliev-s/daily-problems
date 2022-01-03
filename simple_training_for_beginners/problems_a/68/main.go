// https://codeforces.com/problemset/problem/615/A

package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	bulbs := make([]int, m)
	for i, connectedBulbs := 0, 0; i < n; i++ {
		fmt.Scan(&connectedBulbs)
		for j, bulb := 0, 0; j < connectedBulbs; j++ {
			fmt.Scan(&bulb)
			bulbs[bulb-1]++
		}
	}

	allTurnedOn := true
	for i := range bulbs {
		if bulbs[i] == 0 {
			allTurnedOn = false
			break
		}
	}

	if allTurnedOn {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}
