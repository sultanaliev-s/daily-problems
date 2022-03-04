// https://codeforces.com/problemset/problem/1006/C

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

	ar := make([]int64, n)
	for i := range ar {
		fmt.Fscan(in, &ar[i])
	}

	l, r := 0, n-1
	var sumL, sumR, sum int64
	isLeft := true
	for l <= r {
		if isLeft {
			sumL += ar[l]
			l++
		} else {
			sumR += ar[r]
			r--
		}
		if sumL == sumR {
			sum = sumL
			isLeft = true
		} else if (isLeft && sumL > sumR) || (!isLeft && sumL < sumR) {
			isLeft = !isLeft
		}
	}

	fmt.Println(sum)
}
