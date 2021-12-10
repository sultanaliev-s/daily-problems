// https://codeforces.com/problemset/problem/141/A

package main

import (
	"fmt"
	"reflect"
)

func main() {
	var guestName, hostName, heap string
	fmt.Scanln(&guestName)
	fmt.Scanln(&hostName)
	fmt.Scanln(&heap)

	initialLetters := make(map[byte]int)
	heapLetters := make(map[byte]int)

	for i := range guestName {
		initialLetters[guestName[i]]++
	}
	for i := range hostName {
		initialLetters[hostName[i]]++
	}
	for i := range heap {
		heapLetters[heap[i]]++
	}

	areEqual := reflect.DeepEqual(initialLetters, heapLetters)
	if areEqual {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
