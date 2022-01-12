// https://codeforces.com/problemset/problem/1144/B

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(in, &n)

	evens := make([]int, n)
	odds := make([]int, n)
	nEven, nOdd := 0, 0
	for i, x := 0, 0; i < n; i++ {
		fmt.Fscan(in, &x)
		if x%2 == 0 {
			evens[nEven] = x
			nEven++
		} else {
			odds[nOdd] = x
			nOdd++
		}
	}

	sort.Ints(evens)
	sort.Ints(odds)

	res := 0
	if nEven > nOdd {
		res = acc(evens, n-nEven, n-nOdd-1)
	} else {
		res = acc(odds, n-nOdd, n-nEven-1)
	}

	fmt.Println(res)

}

func acc(arr []int, i, n int) (res int) {
	for ; i < n; i++ {
		res += arr[i]
	}
	return
}
