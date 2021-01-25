package main

import (
	"crypto/sha256"
	"fmt"
	"os"

	onhash256 "github.com/sudhir-intc/samples_go/libcrypto/onhash"
)

func main() {

	cmd := os.Args[1]

	fmt.Println("String passed:", cmd)
	h := sha256.New()
	h.Write([]byte(cmd))
	fmt.Printf("Go Crypto Output: %x\n", h.Sum(nil))

	onh := onhash256.New()
	onh.Write([]byte(cmd))
	fmt.Printf("libcrypto Output: %x\n", onh.Sum(nil))

}
