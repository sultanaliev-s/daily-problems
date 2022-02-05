// https://codeforces.com/problemset/problem/158/B

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(in, &n)

	var x, one, two, three, taxi int
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &x)
		switch x {
		case 1:
			one++
		case 2:
			two++
		case 3:
			three++
		case 4:
			taxi++
		}
	}

	taxi += three
	one -= three
	if one < 0 {
		one = 0
	}

	taxi += two / 2
	two %= 2
	one += two * 2
	taxi += one / 4
	one %= 4
	if one > 0 {
		taxi++
	}

	fmt.Println(taxi)
}
