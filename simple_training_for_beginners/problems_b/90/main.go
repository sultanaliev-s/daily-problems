// https://codeforces.com/problemset/problem/136/B

package main

import (
	"fmt"
	"strconv"
)

func toTernary(x int) (res string) {
	r := 0
	for x != 0 {
		r = x % 3
		x /= 3
		res = strconv.Itoa(r) + res
	}
	return res
}

func toDecimal(x string) (res int) {
	for i, j := 0, len(x)-1; j >= 0; i, j = i+1, j-1 {
		res += pow(3, i) * int(x[j]-'0')
	}
	return res
}

func pow(x, y int) int {
	res := 1
	for i := 0; i < y; i++ {
		res *= x
	}
	return res
}

func main() {
	var ad, cd int
	fmt.Scan(&ad, &cd)

	var at, ct string
	at = toTernary(ad)
	ct = toTernary(cd)
	if len(at) < len(ct) {
		t := ""
		for i := len(at); i < len(ct); i++ {
			t += "0"
		}
		at = t + at
	} else {
		t := ""
		for i := len(ct); i < len(at); i++ {
			t += "0"
		}
		ct = t + ct
	}

	var bt string
	for i := range at {
		for j := byte(0); j < 3; j++ {
			if (at[i]-'0'+j)%3 == ct[i]-'0' {
				bt += string(j + '0')
			}
		}
	}

	bd := toDecimal(bt)
	fmt.Println(bd)
}
