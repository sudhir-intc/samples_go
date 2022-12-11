package main

import "fmt"

func main() {
	a := make([]int, 5)
	printSlices(a)
	b := make([]int, 0, 5)
	printSlices(b)

}

func printSlices(s []int) {
	fmt.Printf("len : %d, cap : %d, %v", len(s), cap(s), s)
}
