package main

import "fmt"

func fib() func() int {
	n1 := -1
	n2 := 1
	n3 := 0
	return func() int {
		n3 = n1 + n2
		n1 = n2
		n2 = n3
		return n3
	}
}

func main() {
	f := fib()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
