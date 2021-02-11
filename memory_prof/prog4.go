package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	_ "net/http/pprof"
	"os"
	"runtime"
	"runtime/pprof"
	"sync"
)

type smallStruct struct {
	a, b int64
	c, d float64
}

type heapStruct struct {
	a [64000]int64
}

func main() {

	var input string

	var memprofile = flag.String("memprofile", "", "write memory profile to this file")

	fmt.Println("Performing 4 small (mcache) allocations")
	fmt.Println("Press the Enter Key to continue!")
	fmt.Scanln(&input)
	for i := 0; i < 4; i++ {
		smallAllocation()
		// Print our memory usage at each interval
		PrintMemUsage()

	}
	fmt.Println("Performing 4 heap allocations")
	fmt.Println("Press the Enter Key to continue!")
	fmt.Scanln(&input)
	var overall [][]int
	for i := 0; i < 400; i++ {

		// Allocate memory using make() and append to overall (so it doesn't get
		// garbage collected). This is to create an ever increasing memory usage
		// which we can track. We're just using []int as an example.
		a := make([]int, 0, 100000)
		overall = append(overall, a)
		// Print our memory usage at each interval
		PrintMemUsage()
	}

	fmt.Println("Performing 4 UE Contexts allocations")
	fmt.Println("Press the Enter Key to continue!")
	fmt.Scanln(&input)
	var UEContexts sync.Map
	for i := 0; i < 4; i++ {

		UEContexts.Store(i, NewUEContext())
		// Print our memory usage at each interval
		PrintMemUsage()
	}
	fmt.Println("Performing UE Contexts Deletion and Heap cleanup")
	fmt.Println("Press the Enter Key to continue!")
	fmt.Scanln(&input)
	overall = nil
	for i := 0; i < 4; i++ {
		UEContexts.Delete(i)
	}
	// Print our memory usage at each interval
	PrintMemUsage()
	fmt.Println("Triggering  explicit GC : Press the Enter Key to continue!")
	fmt.Scanln(&input)
	// Force GC to clear up, should see a memory drop
	runtime.GC()

	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.WriteHeapProfile(f)
		f.Close()
		return
	}

}

//go:noinline
func smallAllocation() *smallStruct {
	return &smallStruct{}
}

//go:noinline
func heapAllocation() *heapStruct {
	return &heapStruct{}
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

// AmfUeNgapIDUnspecified the default value for the ID
const (
	AmfUeNgapIDUnspecified int64 = 0xffffffffff
	ranUeNgapID            int64 = 0xffffffffff
)

//SignalingSAState type string
type SignalingSAState string

//SignalingSA states
const (
	//WaitForSignalingSA flag
	WaitForSignalingSA SignalingSAState = "WaitForSignalingSA"
	//SignalingSACreated flag
	SignalingSACreated SignalingSAState = "SignalingSACreated"
	//WaitforSATerminate flag
	WaitforSATerminate SignalingSAState = "WaitforSATerminate"
	//SignalingSATerminated flag
	SignalingSATerminated SignalingSAState = "SignalingSATerminated"
)

//ChildSAState type string
type ChildSAState string

//ChildSA states
const (
	//WaitForChildSA flag
	WaitForChildSA ChildSAState = "WaitForChildSA"
	//ChildSACreated flag
	ChildSACreated ChildSAState = "ChildSACreated"
	//ChildSAFailed flag
	ChildSAFailed ChildSAState = "ChildSAFailed"
	//WaitForChildSATerminate flag
	WaitForChildSATerminate ChildSAState = "WaitForChildSATerminate"
	//ChildSATerminated flag
	ChildSATerminated ChildSAState = "ChildSATerminated"
)

// UEContext is context stored at N3IWF  for each UE
type UEContext struct {
	// RAN UE NGAP ID generated by the N3IWF
	RanUeNgapID int64
	// AMF UE NGAP ID provided by the AMF
	AmfUeNgapID int64
	// UE Location information
	UeLocation Location
	// 32 bytes (256 bits), value is from NGAP IE "Security Key"
	Kn3iwf []uint8
	// NAS TCP Connection information
	TCPConnection net.Conn
	// PDU Session list pduSessionId as key
	//which contains all the successful PDUs
	PduSessionList map[int64]PDUSession
	//Temporary list for PduSession
	UnactivatedPDUSession []PDUSessionUnactivated
	//Temporary list for release PduSession
	UnactivatedReleasePDUSession []PDUSessionUnactivated
	// flag to set state of messages to trigger condition-->
	//WaitForSignalingSA, SignalingSACreated, SignalingSAFailed
	SignalingSAstate SignalingSAState
}

// Location stores the information of the UE location
type Location struct {
	IPAddr     string // UE's wifi address
	PortNumber int32  // UE's IKE port number
}

// PDUSession is Information of a PDU Session as received from the AMF
// It contains a reference to the IP-Sec association as well as the UPF
// GTP Association
type PDUSession struct {
	// PDU Session ID
	PduID int64
	// PDU Type
	PduType int8
	/* Security related information for future */
	QFIList  []uint8
	QosFlows map[int64]*QosFlow // QosFlowIdentifier as key
	//State of ChildSa
	ChildSAstate ChildSAState
}

// PDUSessionUnactivated contains temperary pduinfo
type PDUSessionUnactivated struct {
	PduSess    PDUSession
	CauseValue uint64
	NASPDU     []byte
}

// PDUSessionAggregateMaximumBitRate contains the maximum bit rate info
type PDUSessionAggregateMaximumBitRate struct {
	PDUSessionAggregateMaximumBitRateDL int64
	PDUSessionAggregateMaximumBitRateUL int64
}

// QosFlow contains the QoS Flow info
type QosFlow struct {
	Identifier int64
	Parameters QosFlowLevelQosParameters
}

// QosFlowLevelQosParameters contains the QoS Flow Parameters info
type QosFlowLevelQosParameters struct {
	Placeholder string
}

func NewUEContext() *UEContext {
	ueContext := new(UEContext)
	ueContext.AmfUeNgapID = AmfUeNgapIDUnspecified
	ueContext.RanUeNgapID = ranUeNgapID
	return ueContext
}