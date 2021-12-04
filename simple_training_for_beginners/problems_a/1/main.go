package main

import "fmt"

func main() {
	var nProblems int
	fmt.Scanf("%d\n", &nProblems)

	var nSolvable int = 0
	var petya, vasya, tonya int
	for i := 0; i < nProblems; i++ {
		fmt.Scanf("%d %d %d\n", &petya, &vasya, &tonya)
		if (petya + vasya + tonya) > 1 {
			nSolvable++
		}
	}

	fmt.Println(nSolvable)
}
