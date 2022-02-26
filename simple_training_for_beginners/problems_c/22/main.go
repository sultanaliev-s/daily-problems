// https://codeforces.com/problemset/problem/1133/C

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

	ar := make([]int, n)
	for i := range ar {
		fmt.Fscan(in, &ar[i])
	}

	sort.Ints(ar)

	max := 0
	first := 0
	last := 0
	for first < n && last < n {
		if ar[last]-ar[first] <= 5 {
			last++
		} else {
			if num := last - first; max < num {
				max = num
			}
			first++
		}
	}
	if num := last - first; max < num {
		max = num
	}

	fmt.Println(max)

}
