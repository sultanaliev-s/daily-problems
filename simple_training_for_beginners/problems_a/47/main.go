// https://codeforces.com/problemset/problem/732/A

package main

import "fmt"

func main() {
	var k, r int
	fmt.Scanf("%d %d\n", &k, &r)

	shovels := 0
	for i := 1; i <= 10; i++ {
		priceForShovels := k * i
		if ((priceForShovels - r) % 10) == 0 {
			shovels = i
			break
		} else if (priceForShovels % 10) == 0 {
			shovels = i
			break
		}
	}

	fmt.Println(shovels)
}
