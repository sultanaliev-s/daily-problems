// https://codeforces.com/problemset/problem/520/A

package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	var line string
	fmt.Scan(&n, &line)

	line = strings.ToLower(line)
	var alphabet [26]bool
	for i := 0; i < n; i++ {
		alphabet[line[i]-'a'] = true
	}

	isPangram := true
	for _, v := range alphabet {
		if !v {
			isPangram = false
			break
		}
	}

	if isPangram {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
