package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Monday ?")
	today := time.Now().Weekday()

	switch time.Monday {
	case today + 0:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Tommorrow")
	case today + 2:
		fmt.Println("In 2 days")
	default:
		fmt.Println("Far away")
	}

	switch {
	case today == 0:
		fmt.Println("Tommorrow")
	case today == 1:
		fmt.Println("Today")
	case today == 2:
		fmt.Println("In 2 days")
	default:
		fmt.Println("Far away")
	}
}
