package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {

	var str_slice []string
	var string_map map[string]int

	string_map = make(map[string]int)
	str_slice = strings.Fields(s)

	for _, v := range str_slice {
		string_map[v] += 1
	}
	return string_map
}

func main() {
	fmt.Println(WordCount("This is a sample string of a sample of a sample"))
}
