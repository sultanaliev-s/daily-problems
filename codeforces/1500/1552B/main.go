package main

import (
	"bufio"
	"fmt"
	"os"
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

		ar := make([][]int, n)
		for i := range ar {
			ar[i] = make([]int, 5)
			for j := range ar[i] {
				fmt.Fscan(in, &ar[i][j])
			}
		}

		if n == 1 {
			fmt.Fprintln(out, 1)
			continue
		}

		w := 0
		for i := range ar {
			if better(i, w, ar) {
				w = i
			}
		}

		isSups := true
		for i := range ar {
			if i != w {
				if better(i, w, ar) {
					isSups = false
				}
			}
		}

		if isSups {
			fmt.Fprintln(out, w+1)
		} else {
			fmt.Fprintln(out, -1)
		}
	}
}

func better(i, j int, ar [][]int) bool {
	t := 0
	for k := 0; k < 5; k++ {
		if ar[i][k] < ar[j][k] {
			t++
		}
	}
	return t >= 3
}
