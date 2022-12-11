package main

import "fmt"

func main() {

	primes := [9]int{1, 2, 3, 5, 7, 11, 13, 17, 19}

	for i, v := range primes {
		fmt.Printf("%d +1 : %d\n", i+1, v)
	}
}
