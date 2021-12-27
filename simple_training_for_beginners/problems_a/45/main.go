// https://codeforces.com/problemset/problem/735/A

package main

import "fmt"

func main() {
	var n, k int
	var path string
	fmt.Scanf("%d %d\n", &n, &k)
	fmt.Scanf("%s\n", &path)

	var grasshopper int
	var insect int
	for i, v := range path {
		if v == 'G' {
			grasshopper = i
		} else if v == 'T' {
			insect = i
		}
	}

	if grasshopper > insect {
		k = -k
	}

	canReach := false
	for i := grasshopper; 0 <= i && i < n; i += k {
		if path[i] == 'T' {
			canReach = true
			break
		} else if path[i] == '#' {
			break
		}
	}

	if canReach {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
