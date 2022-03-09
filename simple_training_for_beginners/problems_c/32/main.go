// https://codeforces.com/problemset/problem/670/C

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

	langs := make(map[int]int)
	for i, x := 0, 0; i < n; i++ {
		fmt.Fscan(in, &x)
		langs[x]++
	}

	var m int
	fmt.Fscan(in, &m)

	type film struct {
		filmID int
		dub    int
		sub    int
	}

	films := make([]film, m)
	for i, x := 0, 0; i < m; i++ {
		fmt.Fscan(in, &x)
		films[i].filmID = i
		films[i].dub = langs[x]
	}
	for i, x := 0, 0; i < m; i++ {
		fmt.Fscan(in, &x)
		films[i].sub = langs[x]
	}

	sort.Slice(films, func(i, j int) bool {
		if films[i].dub == films[j].dub {
			return films[i].sub > films[j].sub
		}
		return films[i].dub > films[j].dub
	})

	fmt.Println(films[0].filmID + 1)
}
