package main

import "fmt"

func Sqrt(x float64) float64 {
	z := 1.0
	prev := z
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		if prev == z {
			break
		}
		fmt.Printf("%d. value of z: %f\n", i+1, z)
		prev = z
	}
	return z
}

func main() {
	Sqrt(16)
}
