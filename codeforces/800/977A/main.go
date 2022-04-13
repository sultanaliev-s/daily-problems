package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	for i := 0; i < k; i++ {
		rem := n % 10
		if n != 0 && rem == 0 {
			n /= 10
		} else {
			n--
		}
	}

	fmt.Println(n)
}
