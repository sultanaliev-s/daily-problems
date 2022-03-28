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
	var t int
	fmt.Fscan(in, &t)
	factorials := make([]int64, 0, 30)
	for i := 0; i < cap(factorials); i++ {
		if i == 0 {
			factorials = append(factorials, 1)
		} else {
			factorials = append(factorials, factorials[i-1]*int64(i))
		}
		if factorials[i] > 1_000_000_000_000 {
			break
		}
	}
	for ; t > 0; t-- {
		var n int64
		fmt.Fscan(in, &n)

		ind := 0
		for i := len(factorials) - 1; i >= 0; i-- {
			if factorials[i] <= n {
				ind = i
				break
			}
		}

		answers := make([]int, 0)
		solve(ind, 0, n, factorials, &answers)

		mini := 0
		for i := range answers {
			if answers[mini] > answers[i] {
				mini = i
			}
		}

		res := answers[mini]
		fmt.Fprintln(out, res)
	}
}

func solve(i int, res int, temp int64, facts []int64, answers *[]int) {
	if i == -1 {
		for temp != 0 {
			if temp&1 == 1 {
				res++
			}
			temp >>= 1
		}
		*answers = append(*answers, res)
		return
	}

	if temp-facts[i] >= 0 {
		solve(i-1, res+1, temp-facts[i], facts, answers)
	}

	solve(i-1, res, temp, facts, answers)
}
