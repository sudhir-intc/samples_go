package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])

	list := [6]int{1, 2, 3, 4, 5, 6}
	var big int
	for i := 0; i < len(list)+1; i++ {
		if i > big {
			big = i
		}
	}
	fmt.Println("Biggest num:", big)
}
