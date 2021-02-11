package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"os"
	"strconv"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	log.Println("Main Function....")
	go recur(1, []string{"1"})
	// we need a webserver to get the pprof webserver
	log.Println("Serving on port : 8080")
	log.Println(http.ListenAndServe("localhost:8080", nil))
	os.Exit(1)
}

//recur is a func that intentionally unboundedly increases go routine
// and memory usage and through profiling
// we can see why and where it does so
func recur(i int, strs []string) {
	fmt.Println("Inside recur...")
	time.Sleep(5 * time.Second)
	s1 := append(strs, strconv.Itoa(i))
	s2 := append(s1, strconv.Itoa(i+1))
	go recur(i+1, s1)
	recur(i+2, s2)
}
