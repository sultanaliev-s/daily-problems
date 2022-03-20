// https://codeforces.com/problemset/problem/1370/C

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var t int
	fmt.Fscan(in, &t)
	for ; t > 0; t-- {
		var x int
		fmt.Fscan(in, &x)

		a := true
		if x == 1 {
			a = false
		} else if x%2 == 0 && x > 2 {
			if isPowerOfTwo(x) {
				a = false
			} else if x%4 != 0 && isPrime(x/2) {
				a = false
			}
		}

		if a {
			fmt.Fprintln(out, "Ashishgup")
		} else {
			fmt.Fprintln(out, "FastestFinger")
		}
	}
}

func isPrime(x int) bool {
	lim := int(math.Sqrt(float64(x)))
	for i := 2; i <= lim; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func isPowerOfTwo(x int) bool {
	for x != 1 {
		if x%2 != 0 {
			return false
		}
		x /= 2
	}
	return true
}
