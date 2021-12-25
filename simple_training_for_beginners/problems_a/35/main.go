// https://codeforces.com/problemset/problem/32/B

package main

import "fmt"

func main() {
	var line string
	fmt.Scanln(&line)

	dashPreceded := false
	for i := 0; i < len(line); i++ {
		if !dashPreceded && line[i] == '.' {
			fmt.Print(0)
		} else if dashPreceded && line[i] == '.' {
			dashPreceded = false
			fmt.Print(1)
		} else if dashPreceded {
			dashPreceded = false
			fmt.Print(2)
		} else {
			dashPreceded = true
		}
	}
	fmt.Println()
}
