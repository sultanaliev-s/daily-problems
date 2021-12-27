// https://codeforces.com/problemset/problem/731/A

package main

import "fmt"

func main() {
	var name string
	fmt.Scanf("%s\n", &name)

	rotations := 0
	var prev byte
	for i := range name {
		cur := name[i] - 'a'
		r := cur
		if prev > cur {
			r += 26 - prev
		} else {
			r -= prev
		}
		l := 25 - cur
		if prev < cur {
			l += prev + 1
		} else {
			l -= 25 - prev
		}
		rotations += int(Min(l, r))
		prev = cur
	}

	fmt.Println(rotations)
}

func Min(a, b byte) byte {
	if a > b {
		return b
	}
	return a
}
