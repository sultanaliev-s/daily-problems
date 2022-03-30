package main

import (
	"bufio"
	"fmt"
	"os"
)

type Segment struct {
	L int
	R int
	C int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var tests int
	fmt.Fscan(in, &tests)
	for ; tests > 0; tests-- {
		var n int
		fmt.Fscan(in, &n)

		ar := make([]Segment, n)
		for i := range ar {
			fmt.Fscan(in, &ar[i].L, &ar[i].R, &ar[i].C)
		}

		minSeg := ar[0]
		maxSeg := ar[0]
		bothSeg := ar[0]
		for i := range ar {
			minSeg = Min(minSeg, ar[i])
			maxSeg = Max(maxSeg, ar[i])
			res := 0
			if minSeg == maxSeg {
				res = minSeg.C
			} else {
				res = minSeg.C + maxSeg.C
			}

			if minSeg.L >= ar[i].L && maxSeg.R <= ar[i].R && res >= ar[i].C &&
				bothCheck(bothSeg, ar[i]) {
				bothSeg = ar[i]
			}

			if minSeg.L >= bothSeg.L &&
				maxSeg.R <= bothSeg.R && res > bothSeg.C {
				res = bothSeg.C
			}

			fmt.Fprintln(out, res)
		}
	}
}

func bothCheck(bothSeg, s Segment) bool {
	if bothSeg.L == s.L && bothSeg.R == s.R && bothSeg.C > s.C {
		return true
	}
	if bothSeg.L > s.L && bothSeg.R < s.R ||
		bothSeg.L > s.L && bothSeg.R == s.R ||
		bothSeg.L == s.L && bothSeg.R < s.R {
		return true
	}
	return false
}

func Min(a, b Segment) Segment {
	if a.L < b.L {
		return a
	} else if a.L > b.L {
		return b
	} else {
		if a.C < b.C {
			return a
		}
		return b
	}
}

func Max(a, b Segment) Segment {
	if a.R > b.R {
		return a
	} else if a.R < b.R {
		return b
	} else {
		if a.C < b.C {
			return a
		}
		return b
	}
}
