package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])
	m["Question"] = 1
	fmt.Println("The value:", m["Question"])
	v, ok := m["Not Present"]
	if ok {
		fmt.Println("The value:", v)
	} else {
		fmt.Println("The value not present")
	}
}
