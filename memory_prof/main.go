package main

import (
	"fmt"
	"log"
	_ "net/http/pprof"
	"runtime"
	"time"
)

func main() {

	go func() {
		// Below is an example of using our PrintMemUsage() function
		// Print our starting memory usage (should be around 0mb)
		PrintMemUsage()

		var overall [][]int
		for i := 0; i < 4; i++ {

			// Allocate memory using make() and append to overall (so it doesn't get
			// garbage collected). This is to create an ever increasing memory usage
			// which we can track. We're just using []int as an example.
			a := make([]int, 0, 100000)
			overall = append(overall, a)

			// Print our memory usage at each interval
			PrintMemUsage()
			time.Sleep(time.Second)
		}

		// Clear our memory and print usage, unless the GC has run 'Alloc' will remain the same
		overall = nil
		PrintMemUsage()

		// Force GC to clear up, should see a memory drop
		runtime.GC()
		PrintMemUsage()

	}()

	// we need a webserver to get the pprof webserver
	log.Println("Serving on port : 8080")
	//log.Println(http.ListenAndServe("10.190.212.144:8080", nil))

	fmt.Println("Press the Enter Key to terminate the console screen!")
	var input string
	fmt.Scanln(&input)

}

// PrintMemUsage outputs the current, total and OS memory being used. As well as the number
// of garage collection cycles completed.
func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tMallocs = %v", m.Mallocs)
	fmt.Printf("\tFrees = %v", m.Frees)
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tPauseTotalNs = %v ns", m.PauseTotalNs)
	fmt.Printf("\tNumGC = %v", m.NumGC)
	fmt.Printf("\tForced NumGC = %v\n", m.NumForcedGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
