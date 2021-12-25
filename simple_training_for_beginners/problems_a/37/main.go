// https://codeforces.com/problemset/problem/9/A

package main

import "fmt"

func main() {
	var y, w int
	fmt.Scanf("%d %d\n", &y, &w)

	if y < w {
		y, w = w, y
	}

	winningPossibilities := 6 - y + 1

	switch winningPossibilities {
	case 1:
		fmt.Println("1/6")
	case 2:
		fmt.Println("1/3")
	case 3:
		fmt.Println("1/2")
	case 4:
		fmt.Println("2/3")
	case 5:
		fmt.Println("5/6")
	case 6:
		fmt.Println("1/1")
	default:
		fmt.Println("0/1")
	}
}
