// https://codeforces.com/problemset/problem/894/A

package main

import "fmt"

func main() {
	var line string
	fmt.Scan(&line)

	QAQ := 0
	for i := 0; i < len(line); i++ {
		if line[i] == 'Q' {
			for j := i + 1; j < len(line); j++ {
				if line[j] == 'A' {
					for k := j + 1; k < len(line); k++ {
						if line[k] == 'Q' {
							QAQ++
						}
					}
				}
			}
		}
	}

	fmt.Println(QAQ)
}
