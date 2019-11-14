package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	// Parent context
	parentCtx := context.Background()
	// Create a child context which has value of one second
	ctx, cancel := context.WithTimeout(parentCtx, time.Second)
	// Alternate options :
	// ctx, cancel := context.WithCancel(parentCtx)
	// time.AfterFunc(time.Second, cancel)
	// required to cancel to free resouces created by context time.
	defer cancel()
	//mysleepandprintInCorrect(ctx, 5*time.Second, "Hello")
	mysleepandprintCorrect(ctx, 5*time.Second, "Hello")
}

// mysleepandprintInCorrect incorrect usage of contexts as though parent is
// sending cancel its not handled
func mysleepandprintInCorrect(ctx context.Context, d time.Duration, msg string) {
	time.Sleep(d)
	fmt.Println(msg)
}

// mysleepandprintCorrect correct usage of context
func mysleepandprintCorrect(ctx context.Context, d time.Duration, msg string) {
	select {
	case <-time.After(d):
		fmt.Println(msg)
	case <-ctx.Done():
		log.Print(ctx.Err())
	}
}
