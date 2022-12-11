package main

import (
	"fmt"
)

func main() {
	defer fmt.Printf("\tWorld\n")
	fmt.Printf("Hello")
}
