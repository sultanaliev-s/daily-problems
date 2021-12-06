// https://codeforces.com/problemset/problem/181/A

package main

import "fmt"

type Point struct {
	X int
	Y int
}

func main() {
	var n, m int
	fmt.Scanf("%d %d\n", &n, &m)

	var points []Point
	for i := 0; i < n; i++ {
		var line string
		fmt.Scanln(&line)

		for j := 0; j < m; j++ {
			if line[j] == '*' {
				points = append(points, Point{i, j})
			}
		}
	}

	x := (points[0].X ^ points[1].X ^ points[2].X) + 1
	y := (points[0].Y ^ points[1].Y ^ points[2].Y) + 1
	fmt.Println(x, y)

}
