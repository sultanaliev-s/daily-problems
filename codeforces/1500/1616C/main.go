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

		ar := make([]int, n)
		for i := range ar {
			fmt.Fscan(in, &ar[i])
		}

		if n <= 2 {
			fmt.Fprintln(out, 0)
			continue
		}

		res := n
		for i := 0; i < n; i++ {
			for j := i + 1; j < n; j++ {
				count := 0
				for k := 0; k < n; k++ {
					if (ar[k]-ar[i])*(k-j) != (ar[k]-ar[j])*(k-i) {
						count++
					}
				}

				if count < res {
					res = count
				}
			}
		}

		fmt.Fprintln(out, res)
	}
}
