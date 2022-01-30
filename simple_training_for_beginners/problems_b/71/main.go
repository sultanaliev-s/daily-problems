// https://codeforces.com/problemset/problem/670/B

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n, k int
	fmt.Fscan(in, &n, &k)

	ar := make([]int, n)
	for i := range ar {
		fmt.Fscan(in, &ar[i])
	}

	i, sum := 1, 0
	for sum < k {
		sum += i
		i++
	}
	i--

	ind := i - (sum - k) - 1
	res := ar[ind]

	fmt.Println(res)
}
