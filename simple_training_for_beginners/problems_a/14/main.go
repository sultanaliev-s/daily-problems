// https://codeforces.com/problemset/problem/148/A

package main

import "fmt"

func lcm(a int, b int) int {
	if b > a {
		a, b = b, a
	}

	for i := a; a%b != 0; a += i {
	}

	return a
}

func main() {
	var k, l, m, n, d int
	fmt.Scanf("%d\n%d\n%d\n%d\n%d\n", &k, &l, &m, &n, &d)

	ones := (d / k) + (d / l) + (d / m) + (d / n)

	pairs := (d / lcm(k, l)) + (d / lcm(k, m)) + (d / lcm(k, n)) +
		(d / lcm(l, m)) + (d / lcm(l, n)) + (d / lcm(m, n))

	triples := (d / lcm(k, lcm(l, m))) + (d / lcm(k, lcm(l, n))) +
		(d / lcm(k, lcm(m, n))) + (d / lcm(l, lcm(m, n)))

	quadruples := d / lcm(k, lcm(l, lcm(m, n)))

	result := ones - pairs + triples - quadruples
	fmt.Println(result)
}
