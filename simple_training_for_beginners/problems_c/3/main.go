// https://codeforces.com/problemset/problem/1144/C

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	arr := make([]int, n)
	for i := range arr {
		fmt.Fscan(in, &arr[i])
	}

	inc := make([]int, 0, n/2)
	incMet := make(map[int]bool)
	dec := make([]int, 0, n/2)
	decMet := make(map[int]bool)
	isPossible := true
	for _, v := range arr {
		if !incMet[v] {
			inc = append(inc, v)
			incMet[v] = true
		} else if !decMet[v] {
			dec = append(dec, v)
			decMet[v] = true
		} else {
			isPossible = false
			break
		}
	}

	if !isPossible {
		fmt.Fprintln(out, "NO")
	} else {
		sort.Ints(inc)
		sort.Sort(sort.Reverse(sort.IntSlice(dec)))
		fmt.Fprintln(out, "YES")
		fmt.Fprintln(out, len(inc))
		for i := range inc {
			if i > 0 {
				fmt.Fprint(out, " ")
			}
			fmt.Fprint(out, inc[i])
		}
		fmt.Fprintln(out)
		fmt.Fprintln(out, len(dec))
		for i := range dec {
			if i > 0 {
				fmt.Fprint(out, " ")
			}
			fmt.Fprint(out, dec[i])
		}
		fmt.Fprintln(out)
	}
}
