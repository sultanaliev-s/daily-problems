// https://codeforces.com/problemset/problem/1333/B

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var t int
	fmt.Fscan(in, &t)
	for ; t > 0; t-- {
		var n int
		fmt.Fscan(in, &n)

		a := make([]int, n)
		b := make([]int, n)
		for i := range a {
			fmt.Fscan(in, &a[i])
		}
		for i := range b {
			fmt.Fscan(in, &b[i])
		}

		isPossible := true
		pos, neg := -1, -1
		for i := range a {
			if a[i] == 1 {
				pos = i
				break
			}
		}
		for i := range a {
			if a[i] == -1 {
				neg = i
				break
			}
		}

		for i := n - 1; i >= 0; i-- {
			if b[i] > a[i] {
				if pos == -1 || pos >= i {
					isPossible = false
					break
				}
			} else if b[i] < a[i] {
				if neg == -1 || neg >= i {
					isPossible = false
					break
				}
			}
		}

		if isPossible {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func has(a []int, ind, x int) bool {
	for i := ind - 1; i >= 0; i-- {
		if a[i] == x {
			return true
		}
	}
	return false
}
