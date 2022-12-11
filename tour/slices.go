package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primes[1:] // contains 2,5,7,11,13
	fmt.Println(s)

	s1 := []int{2, 3, 5, 7, 11, 13, 17}
	printSlice(s1)

	s1 = s1[2:8] // will trigger a panic
	printSlice(s1)

}

func printSlice(s []int) {
	fmt.Printf("len : %d, cap : %d %v \n", len(s), cap(s), s)
}
