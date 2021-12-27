// https://codeforces.com/problemset/problem/711/A

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscanf(in, "%d\n", &n)

	rows := make([]string, n)
	row := ""
	hasPlaces := false
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &row)
		if !hasPlaces && row[:2] == "OO" {
			rows[i] = "++" + row[2:]
			hasPlaces = true
		} else if !hasPlaces && row[3:] == "OO" {
			rows[i] = row[:3] + "++"
			hasPlaces = true
		} else {
			rows[i] = row
		}
	}

	if hasPlaces {
		fmt.Println("YES")
		for i := range rows {
			fmt.Println(rows[i])
		}
	} else {
		fmt.Println("NO")
	}
}
