package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
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
	go func() {
		log.Print("Starting the http server 2")
		log.Fatal(http.ListenAndServe("localhost:9080", nil))
	}()
	log.Fatal(http.ListenAndServe("localhost:10080", nil))

}

func handlerWithoutCtx(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handler Started")
	defer log.Printf("Handler Completed")
	time.Sleep(5 * time.Second)
	fmt.Fprintln(w, "Hello")
}

func handlerWithCtx(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	/* Dump the request received */
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}
	log.Println(string(dump))

	defer log.Printf("Handler Completed")
	select {
	case <-time.After(2 * time.Second):
		fmt.Fprintf(w, "Hello Thanks !!!")
	case <-ctx.Done():
		err := ctx.Err()
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
