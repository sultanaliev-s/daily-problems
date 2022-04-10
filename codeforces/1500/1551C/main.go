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
	var test int
	fmt.Fscan(in, &test)
	for ; test > 0; test-- {
		var n int
		fmt.Fscan(in, &n)

		ar := make([]int, n)
		br := make([]int, n)
		cr := make([]int, n)
		dr := make([]int, n)
		er := make([]int, n)
		var s string
		for i := range ar {
			fmt.Fscan(in, &s)
			var ta, tb, tc, td, te, all int
			for j := range s {
				switch s[j] {
				case 'a':
					ta++
				case 'b':
					tb++
				case 'c':
					tc++
				case 'd':
					td++
				case 'e':
					te++
				}
			}
			all = ta + tb + tc + td + te
			ar[i] = ta + ta - all
			br[i] = tb + tb - all
			cr[i] = tc + tc - all
			dr[i] = td + td - all
			er[i] = te + te - all
		}

		sort.Sort(sort.Reverse(sort.IntSlice(ar)))
		sort.Sort(sort.Reverse(sort.IntSlice(br)))
		sort.Sort(sort.Reverse(sort.IntSlice(cr)))
		sort.Sort(sort.Reverse(sort.IntSlice(dr)))
		sort.Sort(sort.Reverse(sort.IntSlice(er)))

		rr := [5]int{
			count(ar),
			count(br),
			count(cr),
			count(dr),
			count(er),
		}

		max := 0
		for i := range rr {
			if rr[i] > rr[max] {
				max = i
			}
		}

		fmt.Fprintln(out, rr[max])
	}
}

func count(ar []int) int {
	c := 0
	for i := range ar {
		if t := c + ar[i]; t < 1 {
			return i
		}
		c += ar[i]
	}
	return len(ar)
}
