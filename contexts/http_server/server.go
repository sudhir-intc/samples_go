package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {

	/*
		_, cancel := context.WithCancel(context.Background())

		osSignals := make(chan os.Signal, 1)
		signal.Notify(osSignals, os.Interrupt, syscall.SIGTERM)
		go func() {
			sig := <-osSignals
			log.Printf("Received signal: %#v", sig)
		}()
	*/
	http.HandleFunc("/", handlerWithCtx)
	log.Print("Starting the http server")
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))

}

func handlerWithoutCtx(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handler Started")
	defer log.Printf("Handler Completed")
	time.Sleep(5 * time.Second)
	fmt.Fprintln(w, "Hello")
}

func handlerWithCtx(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Printf("Handler Started")
	defer log.Printf("Handler Completed")
	select {
	case <-time.After(5 * time.Second):
		fmt.Fprintln(w, "Hello")
	case <-ctx.Done():
		err := ctx.Err()
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
