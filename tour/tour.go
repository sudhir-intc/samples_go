package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	defer demoDefer()

	declarationUsages()
	demoForLoops()
	demoSwitch()
	demoPointers()
	demoStructs()
	demoArrays()
	demoSlices()
	demoRanges()
	demoMaps()
	demoMethods()
}

// Vertex is a structure of X and Y
type Vertex struct {
	X, Y float64
}

// Abs is a method of Vertex
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Scale is a method of Vertex
func (v *Vertex) Scale(f float64) {
	v.X *= f
	v.Y *= f
}

func demoMethods() {

	fmt.Println("Demo of Methods")
	v := Vertex{3, 4}
	fmt.Println(v)
	fmt.Printf("Absolute value: %v\n", v.Abs())
	v.Scale(20)
	fmt.Println(v)
}

func demoMaps() {
	var m = map[string]int{
		"car":  1,
		"bike": 2,
	}
	fmt.Println(m)
	for v := range m {
		fmt.Printf("map[%s] = %d\n", v, m[v])
	}

}

func demoRanges() {
	fmt.Println("Demo of Ranges")
	pow := []int{1, 4, 9, 16, 25}
	for i, v := range pow {
		fmt.Printf("[index,power] => [%d,%d}\n", i, v)
	}
}

func printSlices(s []int) {
	fmt.Printf("Slice %v len=%d, cap=%d\n", s, len(s), cap(s))
}

func demoSlices() {
	var s []int
	fmt.Println("Demo of slices dynamic increase")
	printSlices(s)
	s = append(s, 0)
	printSlices(s)
	s = append(s, 1, 2, 4, 5)
	printSlices(s)
	s = append(s, 6)
	printSlices(s)

	a := make([]int, 10)
	printSlices(a)

	// Double dimension slice
	matrix := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}
	fmt.Println(matrix)

}

func demoArrays() {
	var names [2]string

	fmt.Println("Demo on Arrays and slices")
	names[0] = "sudhir"
	names[1] = "shankar"
	fmt.Println(names)
	fmt.Println(names[0], names[1])

	//primes := [8]int{2, 3, 5, 7, 11, 13, 17, 19} /* Example of array literal */
	primes := [...]int{2, 3, 5, 7, 11, 13, 17, 19} /* Example of array literal allowing the complier to compute the size*/
	//primes := []int{2, 3, 5, 7, 11, 13, 17, 19} /* Example of slice literal */
	var s []int = primes[0:4] /* can be writted as primes[:4] */
	fmt.Printf("Array :%v\n", primes)
	fmt.Printf("Slice: %v\n", s)
	s[0] = 3
	s[1] = 5
	fmt.Printf("Array after slice modification:%v\n", primes)
	fmt.Printf("Slice len:%d, capacity:%d\n", len(s), cap(s))

	var dyn []int = make([]int, 10, 10)
	// or
	// dyn := make([]int, 10, 10)
	fmt.Printf("Dynamic Slice %v len=%d, cap=%d\n", dyn, len(dyn), cap(dyn))

}

func demoStructs() {
	type ListNode struct {
		data int
		next *ListNode
	}
	fmt.Println("Demo of structs using linked list")
	head := ListNode{25, nil}
	currNode := &head

	var newNode *ListNode
	/* Add some nodes nthe list */
	for i := 0; i < 10; i++ {
		newNode = &ListNode{i, nil}
		currNode.next = newNode
		currNode = currNode.next
	}

	currNode = &head
	for currNode != nil {
		fmt.Printf("%d ->", currNode.data)
		currNode = currNode.next
	}
	fmt.Println("nil")
}

func demoPointers() {

	var p *int32
	var num int32 = 0x12345678
	p = &num
	fmt.Printf("Num: 0x%x, char : %v\n", num, *p)
}

// declarationsUsages to be used for demonstration multiple types of declaration
func declarationUsages() {
	var x int              // declared without initialization
	var y int = 42         // declared with initialization
	var a, b int = 45, 100 // declared multiple variables with initialization
	var c = 42             // type ommitted will be inferred
	d := 56                // shorthand only in function scope

	fmt.Printf("Declarations usages: %v %v %v %v %v %v \n", x, y, a, b, c, d)
}

// demoForLoops to be used for demonstration of for loops
func demoForLoops() {
	i := 0
	sum := 0
	fmt.Println("Loops usages")
	// Variant 1
	for i = 0; i < 10; i++ {
		sum += i
	}
	fmt.Printf("Sum1 : %v\n", sum)

	// Variant 2 like a while loop
	sum = 0
	i = 0
	j := 1
	for j == 1 {
		sum += i
		i++
		if i >= 10 {
			break
		}
	}
	fmt.Printf("Sum2 : %v\n", sum)

	// Variant 3 which is forever
	sum = 0
	i = 0
	for {
		sum += i
		i++
		if i >= 10 {
			break
		}
	}
	fmt.Printf("Sum3 : %v\n", sum)
}

func demoSwitch() {
	today := time.Now().Weekday()
	fmt.Printf("Switch case usages Today: %v (%d)\n", today, today)
	switch time.Saturday {
	case today + 0:
		fmt.Println("Saturday is today")
	case today + 1:
		fmt.Println("Saturday is after 1 day")
	case today + 2:
		fmt.Println("Saturday is after 2 days")
	default:
		fmt.Println("Saturday is not near")
	}
}

func demoDefer() {
	fmt.Println("Deferred counting")
	for i := 0; i < 10; i++ {
		defer fmt.Printf("%d ", i)
	}
	fmt.Println("Done")
}
