package main

import (
	"bufio"
	"fmt"
	"os"
)

const MOD int64 = 998_244_353

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var test int
	fmt.Fscan(in, &test)
	for ; test > 0; test-- {
		var n int64
		fmt.Fscan(in, &n)

		ar := make([]int, n)
		for i := range ar {
			fmt.Fscan(in, &ar[i])
		}

		max := 0
		for i := range ar {
			if ar[i] > ar[max] {
				max = i
			}
		}

		c := 0
		k := 0
		for i := range ar {
			if ar[i] == ar[max] {
				c++
			}
			if ar[i] == ar[max]-1 {
				k++
			}
		}

		var fact int64 = 1
		var except int64 = 1
		for i := int64(1); i <= n; i++ {
			fact *= i
			fact %= MOD
			if i != int64(k+1) {
				except *= i
				except %= MOD
			}
		}

		if c == 1 {
			fact = (fact - except + MOD) % MOD
		}

		fmt.Fprintln(out, fact)
	}
}
