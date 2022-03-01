// https://codeforces.com/problemset/problem/1100/C

package main

import (
	"fmt"
	"math"
)

const conv float64 = math.Pi / 180

func main() {
	var n, r int
	fmt.Scan(&n, &r)

	sector := 360.0 / float64(n)
	a := sector / 2
	b := 90 - a
	e := b / 2
	c := 90 - e
	d := 90 - c
	FE := float64(r) * math.Sin(a*conv)
	BE := FE / math.Sin(d*conv)
	BC := BE / (2 * math.Cos(c*conv))

	fmt.Println(BC)
}
