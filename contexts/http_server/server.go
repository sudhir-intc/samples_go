package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"time"
)

func main() {
	http.HandleFunc("/", handlerWithCtx)
	log.Print("Starting the http and http2 server")
	go func() {
		log.Print("Starting the http server 2")
		log.Fatal(http.ListenAndServe(":9080", nil))
	}()

	log.Fatal(http.ListenAndServe(":10080", nil))

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
