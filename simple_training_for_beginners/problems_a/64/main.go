// https://codeforces.com/problemset/problem/658/A

package main

import "fmt"

func main() {
	var n, c int
	fmt.Scan(&n, &c)

	problems := make([]int, n)
	timeToSolve := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&problems[i])
	}
	for i := 0; i < n; i++ {
		fmt.Scan(&timeToSolve[i])
	}

	limScore := 0
	limTime := 0
	radScore := 0
	radTime := 0

	for i, j := 0, n-1; i < n; i, j = i+1, j-1 {
		limTime += timeToSolve[i]
		limScore += Max(0, problems[i]-c*limTime)
		radTime += timeToSolve[j]
		radScore += Max(0, problems[j]-c*radTime)
	}

	if limScore == radScore {
		fmt.Println("Tie")
	} else if limScore > radScore {
		fmt.Println("Limak")
	} else {
		fmt.Println("Radewoosh")
	}

}

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
