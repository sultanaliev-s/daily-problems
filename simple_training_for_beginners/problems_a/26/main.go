// https://codeforces.com/problemset/problem/71/A

package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d\n", &n)

	var line string
	for i := 0; i < n; i++ {
		fmt.Scanln(&line)
		if len(line) > 10 {
			fmt.Printf("%c%d%c\n", line[0], len(line)-2, line[len(line)-1])
		} else {
			fmt.Println(line)
		}
	}
}
